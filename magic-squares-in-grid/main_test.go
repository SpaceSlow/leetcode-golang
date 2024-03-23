package magic_squares_in_grid

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNumMagicSquaresInside(t *testing.T) {
	grid := [][]int{{4, 3, 8, 4}, {9, 5, 1, 9}, {2, 7, 6, 2}}
	expected := 1

	assert.Equal(t, expected, numMagicSquaresInside(grid))
}
