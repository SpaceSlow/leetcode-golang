package main

// https://leetcode.com/problems/happy-number/description/

func isHappy(n int) bool {
	if n == 1 {
		return true
	} else if n < 5 {
		return false
	}

	res := 0
	for n > 0 {
		res += (n % 10) * (n % 10)
		n /= 10
	}
	return isHappy(res)
}

//func isHappy(n int) bool {
//	visited := make(map[int]bool)
//
//	for {
//		if n == 1 {
//			return true
//		}
//		if _, ok := visited[n]; ok {
//			return false
//		}
//		visited[n] = true
//		res := 0
//		for ; n > 0; n /= 10 {
//			res += (n % 10) * (n % 10)
//		}
//		n = res
//	}
//}
