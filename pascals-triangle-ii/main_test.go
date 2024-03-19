package pascals_triangle_ii

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetRow0(t *testing.T) {
	rowIndex := 0
	expected := []int{1}

	assert.Equal(t, expected, getRow(rowIndex))
}

func TestGetRow1(t *testing.T) {
	rowIndex := 1
	expected := []int{1, 1}

	assert.Equal(t, expected, getRow(rowIndex))
}

func TestGetRow3(t *testing.T) {
	rowIndex := 3
	expected := []int{1, 3, 3, 1}

	assert.Equal(t, expected, getRow(rowIndex))
}

func TestGetRow2(t *testing.T) {
	rowIndex := 2
	expected := []int{1, 2, 1}

	assert.Equal(t, expected, getRow(rowIndex))
}
