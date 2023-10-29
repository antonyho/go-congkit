package data_test

import (
	"os"
	"testing"

	"github.com/antonyho/go-cangjie/internal/data"
	"github.com/stretchr/testify/assert"
)

func TestReadTable(t *testing.T) {
	const expectedNumOfEntry = 5

	testTable, err := os.Open("testdata/table.txt")
	if !assert.NoError(t, err) {
		assert.FailNow(t, "failed loading test data", err)
	}
	cangjieTable, err := data.ReadTable(testTable)
	if !assert.NoError(t, err) {
		assert.FailNow(t, "failed parsing table data", err)
	}
	assert.Len(t, cangjieTable, expectedNumOfEntry,
		"result table size '%d' not as expected '%d'",
		len(cangjieTable), expectedNumOfEntry)
}

func TestReadBuiltinTable(t *testing.T) {
	const expectedNumOfEntry = 75012

	cangjieTable, err := data.ReadBuiltinTable()
	if !assert.NoError(t, err) {
		assert.FailNow(t, "failed parsing table data", err)
	}
	assert.Len(t, cangjieTable, expectedNumOfEntry,
		"result table size '%d' not as expected '%d'",
		len(cangjieTable), expectedNumOfEntry)
}
