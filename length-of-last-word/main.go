package main

import "fmt"

// https://leetcode.com/problems/length-of-last-word/

func lengthOfLastWord(s string) int {
	counter := 0
	for i := len(s) - 1; i >= 0; i-- {
		if counter != 0 && s[i] == ' ' {
			return counter
		} else if counter == 0 && s[i] == ' ' {
			continue
		} else {
			counter++
		}
	}
	return counter
}

func main() {
	fmt.Println(lengthOfLastWord("a "))
}
