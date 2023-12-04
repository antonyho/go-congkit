package engine

import (
	"database/sql"
	"errors"
	"fmt"
	"os"
	"unicode/utf8"

	// SQLite3 driver for the engine, the engine uses SQlite3.
	_ "github.com/mattn/go-sqlite3"
)

type CongkitVersion int

const (
	CongkitV3 CongkitVersion = 3
	CongkitV5 CongkitVersion = 5
)

const (
	DatabaseDSNPattern = "file:%s?mode=ro"
)

type Option func(*Engine)

func WithCongkitV3() Option {
	return func(e *Engine) {
		e.CongkitVersion = CongkitV3
	}
}

func WithCongkitV5() Option {
	return func(e *Engine) {
		e.CongkitVersion = CongkitV5
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
	DefaultCongkitVersion = CongkitV5
	DefaultDatabasePath   = "./congkit.db"
)

type Engine struct {
	CongkitVersion
	OutputSimplified bool // Output Simplified Chinese word
	Easy             bool // "Easy" input method mode
	Prediction       bool // Predict word while typing
	dbPath           string
	db               *sql.DB
	query            string
}

func New(options ...Option) *Engine {
	e := &Engine{
		CongkitVersion:   DefaultCongkitVersion,
		OutputSimplified: false,
		dbPath:           DefaultDatabasePath,
	}

	for _, option := range options {
		option(e)
	}

	if _, err := os.Stat(e.dbPath); err != nil && errors.Is(err, os.ErrNotExist) {
		// The Congkit database does not exist, create an in-memory database.
		// This is a constructor. Trying not to return error here.
		e.db, _ = sql.Open("sqlite3", ":memory:")
	} else {
		dsn := fmt.Sprintf(DatabaseDSNPattern, e.dbPath)
		e.db, _ = sql.Open("sqlite3", dsn)
	}

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
	err = e.db.Ping()
	if err != nil {
		return
	}

	results = make([]rune, 0)
	codes := radicals
	if e.Easy {
		if len(radicals) > 1 {
			codes = fmt.Sprintf("%c%%%c", radicals[0], radicals[1])
		}
	} else if e.Prediction {
		codes = fmt.Sprintf("%s%%", radicals)
	}

	rows, err := e.db.Query(e.query, e.CongkitVersion, codes)
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
			e.query = GetSimplifiedCharFromCongkit
		}
	} else {
		if e.Easy {
			e.query = GetCharFromQuick
		} else if e.Prediction {
			e.query = GetCharWithPrediction
		} else {
			e.query = GetCharFromCongkit
		}
	}
}
