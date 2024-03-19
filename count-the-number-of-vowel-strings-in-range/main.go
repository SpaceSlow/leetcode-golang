package count_the_number_of_vowel_strings_in_range

import (
	"strings"
)

// https://leetcode.com/problems/count-the-number-of-vowel-strings-in-range/

func vowelStrings(words []string, left int, right int) int {
	counter := 0

	for i := left; i <= right; i++ {
		if strings.ContainsRune("aeiou", rune(words[i][0])) && strings.ContainsRune("aeiou", rune(words[i][len(words[i])-1])) {
			counter++
		}
	}
	return counter
}
