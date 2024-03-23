package goat_latin

import (
	"strings"
)

// https://leetcode.com/problems/goat-latin/

func toGoatLatin(sentence string) string {
	words := strings.Fields(sentence)

	for i, word := range words {
		words[i] = toGoatLatinWord(word, i+1)
	}

	return strings.Join(words, " ")
}

func toGoatLatinWord(word string, aNumber int) string {
	switch word[0] {
	case 'a':
	case 'e':
	case 'o':
	case 'i':
	case 'u':
	case 'A':
	case 'E':
	case 'O':
	case 'I':
	case 'U':
	default:
		word = word[1:] + string(word[0])
	}
	word += "ma"
	for i := 0; i < aNumber; i++ {
		word += "a"
	}
	return word
}
