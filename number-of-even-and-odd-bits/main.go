package number_of_even_and_odd_bits

// https://leetcode.com/problems/number-of-even-and-odd-bits/description/

//func evenOddBit(n int) []int {
//	even, odd := 0, 0
//
//	for i := 0; n > 0; i++ {
//		if n%2 == 1 {
//			if i%2 == 0 {
//				even++
//			} else {
//				odd++
//			}
//		}
//		n /= 2
//	}
//
//	return []int{even, odd}
//}

func evenOddBit(n int) []int {
	result := make([]int, 2)

	for i := 0; n > 0; i ^= 1 {
		result[i] += n & 1
		n >>= 1
	}

	return result
}
