package main

import "fmt"

// https://leetcode.com/problems/longest-common-prefix/description/

func longestCommonPrefix(strs []string) string {
	firstStr := strs[0]
	minLen := len(firstStr)

	for i := 1; i < len(strs); i++ {
		if len(strs[i]) < minLen {
			minLen = len(strs[i])
		}
		for j := 0; j < minLen; j++ {
			if firstStr[j] != strs[i][j] {
				minLen = j
				break
			}
		}
	}

	return firstStr[0:minLen]
}

func main() {
	strs := []string{
		"flower",
		"flow",
		"flight",
	}
	fmt.Println(longestCommonPrefix(strs))
}
