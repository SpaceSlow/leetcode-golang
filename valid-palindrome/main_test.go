package valid_palindrome

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIsPalindromePanamaPhrase(t *testing.T) {
	str := "A man, a plan, a canal: Panama"
	expected := true

	assert.Equal(t, expected, isPalindrome(str))
}

func TestIsPalindromeRaceACar(t *testing.T) {
	str := "race a car"
	expected := false

	assert.Equal(t, expected, isPalindrome(str))
}

func TestIsPalindromeSpace(t *testing.T) {
	str := " "
	expected := true

	assert.Equal(t, expected, isPalindrome(str))
}
