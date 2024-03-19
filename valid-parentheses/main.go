package main

import "fmt"

// https://leetcode.com/problems/valid-parentheses/

//func isValid(s string) bool {
//	length := len(s)
//	if length%2 != 0 {
//		return false
//	}
//	brackets := map[rune]rune{
//		'}': '{',
//		')': '(',
//		']': '[',
//	}
//	stack := make([]rune, length/2)
//	cursor := 0
//	for i := 0; i < length; i++ {
//		if val, ok := brackets[rune(s[i])]; !ok && cursor < length/2 {
//			stack[cursor] = rune(s[i])
//			cursor++
//		} else if cursor > 0 && stack[cursor-1] == val {
//			cursor--
//		} else {
//			return false
//		}
//	}
//	return cursor == 0
//}

func isValid(s string) bool {
	length := len(s)
	if length%2 != 0 {
		return false
	}
	brackets := map[byte]byte{
		'}': '{',
		')': '(',
		']': '[',
	}

	stack := make([]byte, 0, length/2)
	for i := 0; i < length; i++ {
		switch s[i] {
		case '}', ']', ')':
			{
				if len(stack) > 0 && stack[len(stack)-1] == brackets[s[i]] {
					stack = stack[:len(stack)-1]
				} else {
					return false
				}
			}
		default:
			{
				stack = append(stack, s[i])
			}
		}
	}
	return len(stack) == 0
}

func main() {
	fmt.Println(isValid("(("))
	fmt.Println(isValid("[](){}"))
	fmt.Println(isValid("(()("))
}
