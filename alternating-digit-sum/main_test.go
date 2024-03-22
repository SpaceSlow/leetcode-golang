package alternating_digit_sum

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAlternateDigitSum521(t *testing.T) {
	n := 521
	expected := 4

	assert.Equal(t, expected, alternateDigitSum(n))
}

func TestAlternateDigitSum111(t *testing.T) {
	n := 111
	expected := 1

	assert.Equal(t, expected, alternateDigitSum(n))
}

func TestAlternateDigitSum886996(t *testing.T) {
	n := 886996
	expected := 0

	assert.Equal(t, expected, alternateDigitSum(n))
}
