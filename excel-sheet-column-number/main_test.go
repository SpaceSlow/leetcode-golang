package excel_sheet_column_number

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTitleToNumberA(t *testing.T) {
	column := "A"
	expected := 1

	assert.Equal(t, expected, titleToNumber(column))
}
func TestTitleToNumberAB(t *testing.T) {
	column := "AB"
	expected := 28

	assert.Equal(t, expected, titleToNumber(column))
}

func TestTitleToNumberZY(t *testing.T) {
	column := "ZY"
	expected := 701

	assert.Equal(t, expected, titleToNumber(column))
}
