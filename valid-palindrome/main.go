package valid_palindrome

import (
	"unicode"
)

// https://leetcode.com/problems/valid-palindrome/submissions/

func isPalindrome(s string) bool {
	NL := []*unicode.RangeTable{unicode.Number, unicode.Letter}
	end := len(s) - 1
	for start := 0; start < end+1; start++ {
		s1 := rune(s[start])
		if !unicode.IsOneOf(NL, s1) {
			continue
		}
		s1 = unicode.ToLower(s1)
		for ; end > start; end-- {
			s2 := rune(s[end])
			if unicode.IsOneOf(NL, s2) {
				if s1 != unicode.ToLower(s2) {
					return false
				} else {
					end--
					break
				}
			}
		}
	}
	return true
}
