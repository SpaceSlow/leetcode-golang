package valid_perfect_square

// https://leetcode.com/problems/valid-perfect-square/description/

//func isPerfectSquare(num int) bool {
//	if num == 1 {
//		return true
//	}
//	high := num / 2
//	low := 0
//	for high != low {
//		if high*high == num {
//			return true
//		} else if high*high > num {
//			high = low + (high-low)/2
//		} else {
//			high, low = high*2, high
//		}
//	}
//
//	return false
//}

func isPerfectSquare(num int) bool {
	if num == 1 {
		return true
	}
	high := num / 2
	low := 0
	for low <= high {
		mid := low + (high-low)/2
		if mid*mid == num {
			return true
		} else if mid*mid > num {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	return false
}
