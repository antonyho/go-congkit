package db

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"unicode/utf8"

	"github.com/antonyho/go-congkit/db/models"
	// SQLite3 driver for the engine, the engine uses SQlite3.
	_ "github.com/mattn/go-sqlite3"
)

const (
	CreateCharsTableQuery = `
	CREATE TABLE characters (
		idx INTEGER NOT NULL PRIMARY KEY ASC,
		tc TEXT NOT NULL,
		sc TEXT,
		chinese INTEGER,
		big5 INTEGER,
		hkcsc INTEGER,
		zhuyin INTEGER,
		kanji INTEGER,
		hiragana INTEGER,
		katakana INTEGER,
		punctuation INTEGER,
		symbol INTEGER
	);
	`

	CreateRadicalsTableQuery = `
	CREATE TABLE radicals (
		char_idx INTEGER NOT NULL,
		version INTEGER NOT NULL,
		radical TEXT NOT NULL,
		FOREIGN KEY(char_idx) REFERENCES characters(idx)
	);
	`

	CreateRadicalsIndexQuery = `CREATE INDEX idx_radicals on radicals(version, radical);`

	AddCharsQuery = `
	INSERT INTO characters (
		idx, tc, sc, chinese, big5, hkcsc, zhuyin, kanji, 
		hiragana, katakana, punctuation, symbol
	) 
	VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?);
	`

	AddRadicalsQuery = `
	INSERT INTO radicals (char_idx, version, radical) 
	VALUES (?, ?, ?);
	`
)

// Generate SQLite3 database file from raw data
func Generate(raw [][]string, dbFilePath string) error {
	if _, err := os.Stat(dbFilePath); !errors.Is(err, os.ErrNotExist) {
		log.Println("File already exists at ", dbFilePath)
		log.Println("Remove existing file")
		if err := os.Remove(dbFilePath); err != nil {
			return fmt.Errorf("failed to remove existing db file. %w", err)
		}
	}

	db, err := sql.Open("sqlite3", dbFilePath)
	if err != nil {
		return fmt.Errorf("failed to create db file at %s. %w", dbFilePath, err)
	}
	defer db.Close()

	if _, err := db.Exec(CreateCharsTableQuery); err != nil {
		return fmt.Errorf("error creating 'characters' table. %w", err)
	}
	if _, err := db.Exec(CreateRadicalsTableQuery); err != nil {
		return fmt.Errorf("error creating 'radicals' table. %w", err)
	}
	if _, err := db.Exec(CreateRadicalsIndexQuery); err != nil {
		return fmt.Errorf("error creating index for 'radicals' table. %w", err)
	}

	tx, err := db.Begin()
	if err != nil {
		return fmt.Errorf("error starting db transaction. %w", err)
	}
	addCharStmt, err := tx.Prepare(AddCharsQuery)
	if err != nil {
		return fmt.Errorf("error preparing insert into 'characters' table statement. %w", err)
	}
	defer addCharStmt.Close()
	addRadicalStmt, err := tx.Prepare(AddRadicalsQuery)
	if err != nil {
		return fmt.Errorf("error preparing insert into 'radicals' table statement. %w", err)
	}
	defer addRadicalStmt.Close()

	for rowNum, row := range raw {
		char, radicalSets := convert(rowNum, row)
		if _, err := addCharStmt.Exec(
			char.Idx,
			string(char.Tradition),
			string(char.Simplified),
			char.Chinese,
			char.Big5,
			char.HKSCS,
			char.Zhuyin,
			char.Kanji,
			char.Hiragana,
			char.Katakana,
			char.PunctuationMark,
			char.MiscSymbol,
		); err != nil {
			return fmt.Errorf("error inserting '%c' into 'characters' table. %w", char.Tradition, err)
		}
		for _, radicalSet := range radicalSets {
			if _, err := addRadicalStmt.Exec(
				radicalSet.CharIdx,
				radicalSet.Version,
				radicalSet.Radical,
			); err != nil {
				return fmt.Errorf("error inserting '%c' radical '%s' into 'radicals' table. %w",
					char.Tradition, radicalSet.Radical, err)
			}
		}
	}

	if err := tx.Commit(); err != nil {
		return fmt.Errorf("error committing inserted transactions to db. %w", err)
	}

	return nil
}

func convert(idx int, row []string) (models.Character, []models.RadicalSet) {
	tc, _ := utf8.DecodeRuneInString(row[0])
	var sc rune
	if row[1] != "NA" {
		sc, _ = utf8.DecodeRuneInString(row[1])
	}
	chinese, err := strconv.Atoi(row[2])
	if err != nil {
		log.Printf("Unable to convert rune '%c' column 3 to int", tc)
	}
	big5, err := strconv.Atoi(row[3])
	if err != nil {
		log.Printf("Unable to convert rune '%c' column 4 to int", tc)
	}
	hkcsc, err := strconv.Atoi(row[4])
	if err != nil {
		log.Printf("Unable to convert rune '%c' column 5 to int", tc)
	}
	zhuyin, err := strconv.Atoi(row[5])
	if err != nil {
		log.Printf("Unable to convert rune '%c' column 6 to int", tc)
	}
	kanji, err := strconv.Atoi(row[6])
	if err != nil {
		log.Printf("Unable to convert rune '%c' column 7 to int", tc)
	}
	hiragana, err := strconv.Atoi(row[7])
	if err != nil {
		log.Printf("Unable to convert rune '%c' column 8 to int", tc)
	}
	katakana, err := strconv.Atoi(row[8])
	if err != nil {
		log.Printf("Unable to convert rune '%c' column 9 to int", tc)
	}
	punctuationMark, err := strconv.Atoi(row[9])
	if err != nil {
		log.Printf("Unable to convert rune '%c' column 10 to int", tc)
	}
	miscSymbol, err := strconv.Atoi(row[10])
	if err != nil {
		log.Printf("Unable to convert rune '%c' column 11 to int", tc)
	}

	char := models.Character{
		Idx:             idx,
		Tradition:       tc,
		Simplified:      sc,
		Chinese:         chinese,
		Big5:            big5,
		HKSCS:           hkcsc,
		Zhuyin:          zhuyin,
		Kanji:           kanji,
		Hiragana:        hiragana,
		Katakana:        katakana,
		PunctuationMark: punctuationMark,
		MiscSymbol:      miscSymbol,
	}

	radicalSets := make([]models.RadicalSet, 0)

	if row[11] != "NA" {
		v3radicals := strings.Split(row[11], ",")
		for _, radical := range v3radicals {
			v3radical := models.RadicalSet{
				CharIdx: idx,
				Version: 3,
				Radical: radical,
			}
			radicalSets = append(radicalSets, v3radical)
		}
	}
	if row[12] != "NA" {
		v5radicals := strings.Split(row[12], ",")
		for _, radical := range v5radicals {
			v5radical := models.RadicalSet{
				CharIdx: idx,
				Version: 5,
				Radical: radical,
			}
			radicalSets = append(radicalSets, v5radical)
		}
	}
	if row[13] != "NA" && row[13] != "SPACE" {
		v3radical := models.RadicalSet{
			CharIdx: idx,
			Version: 3,
			Radical: row[13],
		}
		radicalSets = append(radicalSets, v3radical)
		v5radical := models.RadicalSet{
			CharIdx: idx,
			Version: 5,
			Radical: row[13],
		}
		radicalSets = append(radicalSets, v5radical)
	}

	return char, radicalSets
}
