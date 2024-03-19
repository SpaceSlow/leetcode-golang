package excel_sheet_column_title

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTitleToNumber1(t *testing.T) {
	column := 1
	expected := "A"

	assert.Equal(t, expected, convertToTitle(column))
}
func TestTitleToNumber28(t *testing.T) {
	column := 28
	expected := "AB"

	assert.Equal(t, expected, convertToTitle(column))
}

func TestTitleToNumber701(t *testing.T) {
	column := 701
	expected := "ZY"

	assert.Equal(t, expected, convertToTitle(column))
}

func TestTitleToNumber52(t *testing.T) {
	column := 52
	expected := "AZ"

	assert.Equal(t, expected, convertToTitle(column))
}
