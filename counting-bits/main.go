package counting_bits

// https://leetcode.com/problems/counting-bits/

func countBits(n int) []int {
	if n == 0 {
		return []int{0}
	} else if n == 1 {
		return []int{0, 1}
	}
	result := make([]int, n+1)

	j := 0
	result[0], result[1] = 0, 1
	for i := 2; i+j < n+1; j++ {
		if j == i {
			i, j = 2*i, 0
		}
		result[i+j] = 1 + result[j]
	}

	return result
}
