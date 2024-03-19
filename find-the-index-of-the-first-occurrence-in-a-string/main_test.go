package find_the_index_of_the_first_occurrence_in_a_string

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestStrStrTwice(t *testing.T) {
	haystack := "sadbutsad"
	needle := "sad"
	expected := 0

	assert.Equal(t, expected, strStr(haystack, needle))
}
func TestStrStrEnding(t *testing.T) {
	haystack := "adbutsad"
	needle := "sad"
	expected := 5

	assert.Equal(t, expected, strStr(haystack, needle))
}
func TestStrStrHaystackShorterThanNeedle(t *testing.T) {
	haystack := "aaa"
	needle := "aaaa"
	expected := -1

	assert.Equal(t, expected, strStr(haystack, needle))
}
func TestStrStrMississippi(t *testing.T) {
	haystack := "mississippi"
	needle := "issip"
	expected := 4

	assert.Equal(t, expected, strStr(haystack, needle))
}

func TestStrStrNonOccurrences(t *testing.T) {
	haystack := "leetcode"
	needle := "leeto"
	expected := -1

	assert.Equal(t, expected, strStr(haystack, needle))
}
