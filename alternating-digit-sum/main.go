package alternating_digit_sum

// https://leetcode.com/problems/alternating-digit-sum/description/

//func alternateDigitSum(n int) int {
//	digits := make([]int, 0)
//
//	for i := 0; n > 0; i++ {
//		digits = append(digits, n%10)
//		n /= 10
//	}
//
//	alternateSum := 0
//	multiplier := 1
//	for i := len(digits) - 1; i >= 0; i-- {
//		alternateSum += multiplier * digits[i]
//		multiplier *= -1
//	}
//
//	return alternateSum
//}

func alternateDigitSum(n int) int {
	alternateSum := 0
	multiplier := 1
	for ; n > 0; n /= 10 {
		alternateSum += multiplier * n % 10
		multiplier *= -1
	}

	if multiplier == 1 {
		return -alternateSum
	}

	return alternateSum
}
