package ugly_number

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIsUgly6(t *testing.T) {
	num := 6
	expected := true

	assert.Equal(t, isUgly(num), expected)
}

func TestIsUgly1(t *testing.T) {
	num := 1
	expected := true

	assert.Equal(t, isUgly(num), expected)
}

func TestIsUgly14(t *testing.T) {
	num := 14
	expected := false

	assert.Equal(t, isUgly(num), expected)
}

func TestIsUgly0(t *testing.T) {
	num := 0
	expected := false

	assert.Equal(t, isUgly(num), expected)
}

func TestIsUglyMinus2147483648(t *testing.T) {
	num := -2147483648
	expected := false

	assert.Equal(t, isUgly(num), expected)
}
