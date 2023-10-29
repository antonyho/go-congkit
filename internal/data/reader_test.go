package data_test

import (
	"os"
	"testing"

	"github.com/antonyho/go-cangjie/data"
	"github.com/stretchr/testify/assert"
)

func TestReadTable(t *testing.T) {
	testTable, err := os.ReadFile("testdata/table.txt")
	if !assert.NoError(t, err) {
		assert.FailNow(t, "failed loading test data", err)
	}
	cangjieTable, err := data.ReadTable(testTable)
	if !assert.NoError(t, err) {
		assert.FailNow(t, "failed parsing table data", err)
	}
	assert.Len(t, cangjieTable, 5)
}

func TestReadBuiltinTable(t *testing.T) {
	cangjieTable, err := data.ReadBuiltinTable()
	if !assert.NoError(t, err) {
		assert.FailNow(t, "failed parsing table data", err)
	}
	assert.Len(t, cangjieTable, 75017)
}
