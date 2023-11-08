package engine

import (
	"database/sql"

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
	}
}

func WithPrediction() Option {
	return func(e *Engine) {
		e.Prediction = true
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
	Prediction       bool // Predict word with given radicals
	dbPath           string
	db               *sql.DB
}

func New(options ...Option) (*Engine, error) {
	e := &Engine{
		CangjieVersion:   DefaultCangjieVersion,
		OutputSimplified: false,
		dbPath:           DefaultDatabasePath,
	}

	for _, option := range options {
		option(e)
	}

	db, err := sql.Open("sqlite3", e.dbPath)
	if err != nil {
		return nil, err
	}
	e.db = db

	return e, nil
}

func (e *Engine) Set(options ...Option) {
	for _, option := range options {
		option(e)
	}
}

func (e *Engine) Compose(radicals string) (results []rune, err error) {
	results = make([]rune, 0)
	return
}
