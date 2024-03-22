package number_of_beautiful_pairs

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCountBeatifulPairsSingleDigits(t *testing.T) {
	nums := []int{2, 5, 1, 4}
	expected := 5

	assert.Equal(t, expected, countBeautifulPairs(nums))
}

func TestCountBeatifulPairsTwoDigitNumbers(t *testing.T) {
	nums := []int{11, 21, 12}
	expected := 2

	assert.Equal(t, expected, countBeautifulPairs(nums))
}
