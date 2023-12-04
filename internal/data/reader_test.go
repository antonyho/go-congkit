package data_test

import (
	"embed"
	"testing"

	"github.com/antonyho/go-congkit/internal/data"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

//go:embed testdata/table.txt
var testdataCongkitTable embed.FS

func TestReadTable(t *testing.T) {
	const expectedNumOfEntry = 5

	testTable, err := testdataCongkitTable.Open("testdata/table.txt")
	require.NoError(t, err, "failed loading test data")
	congkitTable, err := data.ReadTable(testTable)
	require.NoError(t, err, "failed parsing table data")
	assert.Len(t, congkitTable, expectedNumOfEntry,
		"result table size '%d' not as expected '%d'",
		len(congkitTable), expectedNumOfEntry)
}

func TestReadBuiltinTable(t *testing.T) {
	const expectedNumOfEntry = 75012

	congkitTable, err := data.ReadBuiltinTable()
	require.NoError(t, err, "failed parsing table data")
	assert.Len(t, congkitTable, expectedNumOfEntry,
		"result table size '%d' not as expected '%d'",
		len(congkitTable), expectedNumOfEntry)
}
