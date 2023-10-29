package models

type Character struct {
	Idx             int
	Tradition       rune
	Simplified      rune
	Chinese         int
	Big5            int
	HKSCS           int
	Zhuyin          int
	Kanji           int
	Hiragana        int
	Katakana        int
	PunctuationMark int
	MiscSymbol      int
}
