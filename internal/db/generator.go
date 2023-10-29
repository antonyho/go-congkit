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
	CangjieV3Code       string
	CangjieV5Code       string
	ShortCut            string
	SortOrder           uint16
}
