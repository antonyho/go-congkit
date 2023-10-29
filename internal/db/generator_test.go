package db_test

import (
	"database/sql"
	"os"
	"path"
	"testing"

	"github.com/antonyho/go-cangjie/internal/data"
	"github.com/antonyho/go-cangjie/internal/db"
	_ "github.com/mattn/go-sqlite3"
	"github.com/stretchr/testify/assert"
)

const (
	CountCharsQuery = `SELECT COUNT(ALL) FROM characters;`

	CountRadicalsQuery = `SELECT COUNT(ALL) FROM radicals;`
)

func TestGenerate(t *testing.T) {
	cangjieTable := loadTestTableData(t)

	tempDbFile := path.Join(t.TempDir(), "test.db")
	err := db.Generate(cangjieTable, tempDbFile)
	if !assert.NoError(t, err) {
		assert.FailNow(t, "failed generating database", err)
	}

	assert.FileExists(t, tempDbFile, "db file was not created")

	fileInfo, err := os.Stat(tempDbFile)
	if !assert.NoError(t, err) {
		assert.FailNow(t, "unable to get file info of db file %s", tempDbFile)
	}

	assert.Greater(t, fileInfo.Size(), int64(0), "generated db file size is %d", fileInfo.Size())

	db := openDb(t, tempDbFile)
	defer db.Close()

	result := db.QueryRow(CountCharsQuery)
	var rowCount int
	err = result.Scan(&rowCount)
	if !assert.NoError(t, err) {
		assert.Fail(t, "failed querying 'characters' table row count. %v", err)
	}
	assert.Equal(t, 5, rowCount)

	result = db.QueryRow(CountRadicalsQuery)
	err = result.Scan(&rowCount)
	if !assert.NoError(t, err) {
		assert.Fail(t, "failed querying 'radicals' table row count. %v", err)
	}
	assert.Equal(t, 10, rowCount)
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

func openDb(t *testing.T, dbFilePath string) *sql.DB {
	db, err := sql.Open("sqlite3", dbFilePath)
	if !assert.NoError(t, err) {
		assert.FailNow(t, "failed opening test db")
	}

	return db
}
