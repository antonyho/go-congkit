package data_test

import (
	"embed"
	"testing"

	"github.com/antonyho/go-cangjie/internal/data"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

//go:embed testdata/table.txt
var testdataCangjieTable embed.FS

func TestReadTable(t *testing.T) {
	const expectedNumOfEntry = 5

	testTable, err := testdataCangjieTable.Open("testdata/table.txt")
	require.NoError(t, err, "failed loading test data")
	cangjieTable, err := data.ReadTable(testTable)
	require.NoError(t, err, "failed parsing table data")
	assert.Len(t, cangjieTable, expectedNumOfEntry,
		"result table size '%d' not as expected '%d'",
		len(cangjieTable), expectedNumOfEntry)
}

func TestReadBuiltinTable(t *testing.T) {
	const expectedNumOfEntry = 75012

	cangjieTable, err := data.ReadBuiltinTable()
	require.NoError(t, err, "failed parsing table data")
	assert.Len(t, cangjieTable, expectedNumOfEntry,
		"result table size '%d' not as expected '%d'",
		len(cangjieTable), expectedNumOfEntry)
}
