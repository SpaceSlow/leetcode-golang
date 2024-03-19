package largest_3_same_digit_number_in_string

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLargestGoodIntegerTwoDistinctGoodInteger(t *testing.T) {
	num := "6777133339"
	expected := "777"

	assert.Equal(t, expected, largestGoodInteger(num))
}

func TestLargestGoodIntegerOnlyOneGoodInteger(t *testing.T) {
	num := "2300019"
	expected := "000"

	assert.Equal(t, expected, largestGoodInteger(num))
}

func TestLargestGoodIntegerNoGoodIntegers(t *testing.T) {
	num := "42352338"
	expected := ""

	assert.Equal(t, expected, largestGoodInteger(num))
}
