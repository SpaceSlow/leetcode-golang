package main

import "fmt"

// https://leetcode.com/problems/roman-to-integer/description/

//Symbol       Value
//I             1
//V             5
//X             10
//L             50
//C             100
//D             500
//M             1000

func romanToInt(s string) int {

	romanToIntMap := map[rune]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}

	res, cur := 0, 0
	for index, symbol := range s {
		if index+1 == len(s) {
			res += romanToIntMap[symbol]
			return res
		}
		switch symbol {
		case 'I':
			{
				if s[index+1] == 'V' || s[index+1] == 'X' {
					cur = -1
				} else {
					cur = 1
				}
			}
		case 'X':
			{
				if s[index+1] == 'L' || s[index+1] == 'C' {
					cur = -10
				} else {
					cur = 10
				}
			}
		case 'C':
			{
				if s[index+1] == 'D' || s[index+1] == 'M' {
					cur = -100
				} else {
					cur = 100
				}
			}
		default:
			{
				cur = romanToIntMap[symbol]
			}
		}
		res += cur
	}
	return 0
}

func main() {
	fmt.Println(romanToInt("IX"))
}
