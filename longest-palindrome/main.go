package longest_palindrome

// https://leetcode.com/problems/longest-palindrome/

func longestPalindrome(s string) int {
	freqs := make(map[rune]int)

	for _, letter := range s {
		freqs[letter]++
	}

	maxPalindromeLength := 0
	for _, value := range freqs {
		maxPalindromeLength += value / 2 * 2
	}
	if maxPalindromeLength < len(s) {
		maxPalindromeLength++
	}
	return maxPalindromeLength
}
