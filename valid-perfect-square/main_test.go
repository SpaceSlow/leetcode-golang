package valid_perfect_square

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIsPerfectSquare1(t *testing.T) {
	num := 1
	expected := true

	assert.Equal(t, expected, isPerfectSquare(num))
}

func TestIsPerfectSquare100(t *testing.T) {
	num := 100
	expected := true

	assert.Equal(t, expected, isPerfectSquare(num))
}

func TestIsPerfectSquare2(t *testing.T) {
	num := 2
	expected := false

	assert.Equal(t, expected, isPerfectSquare(num))
}

func TestIsPerfectSquare25(t *testing.T) {
	num := 25
	expected := true

	assert.Equal(t, expected, isPerfectSquare(num))
}

func TestIsPerfectSquare99(t *testing.T) {
	num := 99
	expected := false

	assert.Equal(t, expected, isPerfectSquare(num))
}
