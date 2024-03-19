package power_of_two

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIsPowerOfTwo1(t *testing.T) {
	num := 1
	expected := true

	assert.Equal(t, isPowerOfTwo(num), expected)
}

func TestIsPowerOfTwo16(t *testing.T) {
	num := 16
	expected := true

	assert.Equal(t, isPowerOfTwo(num), expected)
}

func TestIsPowerOfTwo3(t *testing.T) {
	num := 3
	expected := false

	assert.Equal(t, isPowerOfTwo(num), expected)
}

func TestIsPowerOfTwoMinus16(t *testing.T) {
	num := -16
	expected := false

	assert.Equal(t, isPowerOfTwo(num), expected)
}
