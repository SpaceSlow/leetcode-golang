package count_the_number_of_vowel_strings_in_range

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestVowelStrings1(t *testing.T) {
	words := []string{"are", "amy", "u"}
	left := 0
	right := 2
	expected := 2

	assert.Equal(t, expected, vowelStrings(words, left, right))
}

func TestVowelStrings2(t *testing.T) {
	words := []string{"hey", "aeo", "mu", "ooo", "artro"}
	left := 1
	right := 4
	expected := 3

	assert.Equal(t, expected, vowelStrings(words, left, right))
}
