package db

type Entry struct {
	Character           string
	SimplifiedCharacter string
	Chinese             bool
	Big5                bool
	HKSCS               bool
	Zhuyin              bool
	Kanji               bool
	Hiragana            bool
	Katakana            bool
	PunctuationMark     bool
	MiscSymbol          bool
	CangjieV3Radicals   string
	CangjieV5Radicals   string
	ShortCut            string
	SortOrder           uint16
}

const (
	CreateCharsTableQuery = `
	CREATE TABLE characters (
		idx INTEGER NOT NULL PRIMARY KEY ASC,
		tc TEXT NOT NULL UNIQUE,
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
		radicals TEXT NOT NULL,
		FOREIGN KEY(char_idx) REFERENCES characters(idx)
	);
	`

	AddCharsQuery = `
	INSERT INTO characters VALUES ();
	`

	AddRadicalsQuery = `
	INSERT INTO radicals VALUES ();
	`
)

// Generate SQLite3 database file from raw data
func Generate(raw [][]string, dbFilePath string) error {
	return nil
}
