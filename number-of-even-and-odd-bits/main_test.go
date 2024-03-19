package number_of_even_and_odd_bits

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestEvenOddBit17(t *testing.T) {
	n := 17
	expected := []int{2, 0}

	assert.Equal(t, expected, evenOddBit(n))
}

func TestEvenOddBit2(t *testing.T) {
	n := 2
	expected := []int{0, 1}

	assert.Equal(t, expected, evenOddBit(n))
}
