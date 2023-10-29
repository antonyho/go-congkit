package db_test

import (
	"os"
	"path"
	"testing"

	"github.com/antonyho/go-cangjie/internal/data"
	"github.com/antonyho/go-cangjie/internal/db"
	"github.com/stretchr/testify/assert"
)

func TestGenerate(t *testing.T) {
	cangjieTable := loadTestTableData(t)

	tempDbFile := path.Join(t.TempDir(), "test.db")
	err := db.Generate(cangjieTable, tempDbFile)
	if !assert.NoError(t, err) {
		assert.FailNow(t, "failed generating database", err)
	}

	// TODO - Verify the db
}

func loadTestTableData(t *testing.T) [][]string {
	testTable, err := os.Open("testdata/table.txt")
	if !assert.NoError(t, err) {
		assert.FailNow(t, "failed loading test data", err)
	}
	cangjieTable, err := data.ReadTable(testTable)
	if !assert.NoError(t, err) {
		assert.FailNow(t, "failed parsing table data", err)
	}

	return cangjieTable
}
