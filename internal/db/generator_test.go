package db_test

import (
	"database/sql"
	"embed"
	"os"
	"path"
	"testing"

	"github.com/antonyho/go-cangjie/internal/data"
	"github.com/antonyho/go-cangjie/internal/db"
	_ "github.com/mattn/go-sqlite3"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const (
	CountCharsQuery = `SELECT COUNT(ALL) FROM characters;`

	CountRadicalsQuery = `SELECT COUNT(ALL) FROM radicals;`
)

//go:embed testdata/table.txt
var testdataCangjieTable embed.FS

func TestGenerate(t *testing.T) {
	cangjieTable := loadTestTableData(t)

	tempDbFile := path.Join(t.TempDir(), "test.db")
	err := db.Generate(cangjieTable, tempDbFile)
	require.NoError(t, err, "failed generating database")

	assert.FileExists(t, tempDbFile, "db file was not created")

	fileInfo, err := os.Stat(tempDbFile)
	require.NoError(t, err, "unable to get file info of db file %s", tempDbFile)

	assert.Greater(t, fileInfo.Size(), int64(0), "generated db file size is %d", fileInfo.Size())

	db := openDb(t, tempDbFile)
	defer db.Close()

	result := db.QueryRow(CountCharsQuery)
	var rowCount int
	err = result.Scan(&rowCount)
	assert.NoError(t, err, "failed querying 'characters' table row count.")
	assert.Equal(t, 5, rowCount)

	result = db.QueryRow(CountRadicalsQuery)
	err = result.Scan(&rowCount)
	assert.NoError(t, err, "failed querying 'radicals' table row count.")
	assert.Equal(t, 10, rowCount)
}

func loadTestTableData(t *testing.T) [][]string {
	testTable, err := testdataCangjieTable.Open("testdata/table.txt")
	require.NoError(t, err, "failed loading test data")
	cangjieTable, err := data.ReadTable(testTable)
	require.NoError(t, err, "failed parsing table data")

	return cangjieTable
}

func openDb(t *testing.T, dbFilePath string) *sql.DB {
	db, err := sql.Open("sqlite3", dbFilePath)
	require.NoError(t, err, "failed opening test db")

	return db
}
