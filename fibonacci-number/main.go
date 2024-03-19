package fibonacci_number

// https://leetcode.com/problems/fibonacci-number/

func fib(n int) int {
	if n < 2 {
		return n
	}
	f0, f1 := 0, 1

	for i := 2; i <= n; i++ {
		f0, f1 = f1, f0+f1
	}

	return f1
}
