package ugly_number

// https://leetcode.com/problems/ugly-number/description/

func isUgly(n int) bool {
	if n <= 0 {
		return false
	}

	for _, i := range []int{2, 3, 5} {
		for n%i == 0 {
			n /= i
		}
	}

	return n == 1
}
