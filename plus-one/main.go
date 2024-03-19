package main

import "fmt"

func plusOne(digits []int) []int {
	for i := len(digits) - 1; i >= 0; i-- {
		if digits[i] != 9 {
			digits[i]++
			break
		} else if digits[i] == 9 && i != 0 {
			digits[i] = 0
		} else {
			// когда digits[i] == 9 && i == 0
			digits[i] = 0
			digits = append([]int{1}, digits...)
		}
	}
	return digits
}

func main() {
	fmt.Println(plusOne([]int{9, 9}))
}
