package fibonacci_number

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFib2(t *testing.T) {
	n := 2
	expected := 1

	assert.Equal(t, expected, fib(n))
}

func TestFib3(t *testing.T) {
	n := 3
	expected := 2

	assert.Equal(t, expected, fib(n))
}

func TestFib4(t *testing.T) {
	n := 4
	expected := 3

	assert.Equal(t, expected, fib(n))
}
