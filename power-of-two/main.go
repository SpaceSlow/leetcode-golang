package power_of_two

// https://leetcode.com/problems/power-of-two/description/

import (
	"strconv"
	"strings"
)

func isPowerOfTwo(n int) bool {
	if n <= 0 {
		return false
	}
	s := strconv.FormatInt(int64(n), 2)
	return strings.Count(s, "1") == 1
}
