package toeplitz_matrix

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIsToeplitzMatrixToeplitz3x4(t *testing.T) {
	matrix := [][]int{{1, 2, 3, 4}, {5, 1, 2, 3}, {9, 5, 1, 2}}
	expected := true

	assert.Equal(t, expected, isToeplitzMatrix(matrix))
}

func TestIsToeplitzMatrixNotToeplitz2x2(t *testing.T) {
	matrix := [][]int{{1, 2}, {2, 2}}
	expected := false

	assert.Equal(t, expected, isToeplitzMatrix(matrix))
}

func TestIsToeplitzMatrixToeplitz1x1(t *testing.T) {
	matrix := [][]int{{1}}
	expected := true

	assert.Equal(t, expected, isToeplitzMatrix(matrix))
}
