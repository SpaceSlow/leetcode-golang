package pascals_triangle

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGenerate5(t *testing.T) {
	numRows := 5
	expected := [][]int{
		{1},
		{1, 1},
		{1, 2, 1},
		{1, 3, 3, 1},
		{1, 4, 6, 4, 1},
	}
	assert.Equal(t, expected, generate(numRows))
}

func TestGenerate1(t *testing.T) {
	numRows := 1
	expected := [][]int{
		{1},
	}
	assert.Equal(t, expected, generate(numRows))
}
