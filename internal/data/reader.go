package data

import (
	"bufio"
	"embed"
	"errors"
	"fmt"
	"io/fs"
	"strings"
)

// Errors on parsing raw data from embed data file
var (
	ErrMalformEntry = errors.New("data: malform entry")
	ErrCommentLine  = errors.New("data: comment line")
	ErrEmptyLine    = errors.New("data: empty line")
)

//go:embed assets/table.txt
var builtinCongkitTable embed.FS

func ReadTable(congkitTableContent fs.File) ([][]string, error) {
	table := make([][]string, 0)
	scanner := bufio.NewScanner(congkitTableContent)
	for lineNum := 0; scanner.Scan(); lineNum++ {
		entry, err := readRaw(scanner.Text())
		if err != nil {
			switch err {
			case ErrCommentLine, ErrEmptyLine:
				continue
			default:
				return nil, fmt.Errorf("line %d: %w", lineNum, err)
			}
		}
		table = append(table, entry)
	}

	return table, nil
}

func ReadBuiltinTable() ([][]string, error) {
	file, err := builtinCongkitTable.Open("assets/table.txt")
	if err != nil {
		return nil, err
	}
	return ReadTable(file)
}

// Read the line of raw data from the embed Congkit radicals table.
// Refer to the readme file in `assets/` directory for the data format.
func readRaw(line string) ([]string, error) {
	trimmedLine := strings.TrimSpace(line)
	if len(trimmedLine) == 0 {
		return nil, ErrEmptyLine
	}
	if trimmedLine[0] == '#' {
		return nil, ErrCommentLine
	}

	fields := strings.Split(line, " ")
	if len(fields) < 15 {
		return nil, ErrMalformEntry
	}

	return fields, nil
}
