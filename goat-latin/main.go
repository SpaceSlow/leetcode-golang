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
	s := strings.Builder{}
	switch word[0] {
	case 'a':
		fallthrough
	case 'e':
		fallthrough
	case 'o':
		fallthrough
	case 'i':
		fallthrough
	case 'u':
		fallthrough
	case 'A':
		fallthrough
	case 'E':
		fallthrough
	case 'O':
		fallthrough
	case 'I':
		fallthrough
	case 'U':
		s.WriteString(word)
	default:
		{
			s.WriteString(word[1:])
			s.WriteByte(word[0])
		}
	}
	s.WriteString("ma")
	for i := 0; i < aNumber; i++ {
		s.WriteRune('a')
	}
	return s.String()
}
