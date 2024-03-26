package counting_bits

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCountBits2(t *testing.T) {
	n := 2
	expected := []int{0, 1, 1}

	assert.Equal(t, expected, countBits(n))
}

func TestCountBits5(t *testing.T) {
	n := 5
	expected := []int{0, 1, 1, 2, 1, 2}

	assert.Equal(t, expected, countBits(n))
}

func TestCountBits8(t *testing.T) {
	n := 8
	expected := []int{0, 1, 1, 2, 1, 2, 2, 3, 1}

	assert.Equal(t, expected, countBits(n))
}
