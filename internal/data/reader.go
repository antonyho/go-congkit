package data

import (
	"bytes"
	_ "embed"
	"encoding/csv"
)

//go:embed assets/table.txt
var builtinCangjieTable []byte

func ReadTable(cangjieTableContent []byte) ([][]string, error) {
	reader := csv.NewReader(bytes.NewReader(cangjieTableContent))
	reader.Comma = ' '
	reader.Comment = '#'
	reader.LazyQuotes = true
	return reader.ReadAll()
}

func ReadBuiltinTable() ([][]string, error) {
	return ReadTable(builtinCangjieTable)
}
