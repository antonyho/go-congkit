package engine

import (
	"database/sql"
	"errors"
	"fmt"
	"unicode/utf8"

	// SQLite3 driver for the engine, the engine uses SQlite3.
	_ "github.com/mattn/go-sqlite3"
)

type CangjieVersion int

const (
	CangjieV3 CangjieVersion = 3
	CangjieV5 CangjieVersion = 5
)

type Option func(*Engine)

func WithCangjieV3() Option {
	return func(e *Engine) {
		e.CangjieVersion = CangjieV3
	}
}

func WithCangjieV5() Option {
	return func(e *Engine) {
		e.CangjieVersion = CangjieV5
	}
}

func WithSimplified() Option {
	return func(e *Engine) {
		e.OutputSimplified = true
	}
}

func WithEasy() Option {
	return func(e *Engine) {
		e.Easy = true
		e.Prediction = false
	}
}

func WithPrediction() Option {
	return func(e *Engine) {
		e.Prediction = true
		e.Easy = false
	}
}

func WithDatabase(path string) Option {
	return func(e *Engine) {
		e.dbPath = path
	}
}

// Engine defaults
const (
	DefaultCangjieVersion = CangjieV5
	DefaultDatabasePath   = "./cangjie.db"
)

type Engine struct {
	CangjieVersion
	OutputSimplified bool // Output Simplified Chinese word
	Easy             bool // "Easy" input method mode
	Prediction       bool // Predict word while typing
	dbPath           string
	db               *sql.DB
	query            string
}

func New(options ...Option) *Engine {
	e := &Engine{
		CangjieVersion:   DefaultCangjieVersion,
		OutputSimplified: false,
		dbPath:           DefaultDatabasePath,
	}

	for _, option := range options {
		option(e)
	}

	db, _ := sql.Open("sqlite3", e.dbPath)
	e.db = db

	e.determineQuery()

	return e
}

func (e *Engine) Set(options ...Option) {
	for _, option := range options {
		option(e)
	}

	e.determineQuery()
}

func (e *Engine) Encode(radicals string) (results []rune, err error) {
	results = make([]rune, 0)

	codes := radicals
	if e.Easy {
		if len(radicals) > 1 {
			codes = fmt.Sprintf("%c%%%c", radicals[0], radicals[1])
		}
	} else if e.Prediction {
		codes = fmt.Sprintf("%s%%", radicals)
	}

	rows, err := e.db.Query(e.query, e.CangjieVersion, codes)
	if err != nil {
		return
	}
	defer rows.Close()

	for rows.Next() {
		var s string
		scanErr := rows.Scan(&s)
		if scanErr != nil {
			err = errors.Join(scanErr, err)
		}
		char, _ := utf8.DecodeRuneInString(s)
		if char != 0 {
			results = append(results, char)
		}
	}
	rowsErr := rows.Err()
	err = errors.Join(rowsErr, err)

	return
}

func (e *Engine) Close() error {
	return e.db.Close()
}

func (e *Engine) determineQuery() {
	if e.OutputSimplified {
		if e.Easy {
			e.query = GetSimplifiedCharFromQuick
		} else if e.Prediction {
			e.query = GetSimplifiedCharWithPrediction
		} else {
			e.query = GetSimplifiedCharFromCangjie
		}
	} else {
		if e.Easy {
			e.query = GetCharFromQuick
		} else if e.Prediction {
			e.query = GetCharWithPrediction
		} else {
			e.query = GetCharFromCangjie
		}
	}
}
