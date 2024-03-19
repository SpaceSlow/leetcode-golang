package add_digits

// https://leetcode.com/problems/add-digits/

//func addDigits(num int) int {
//	if num < 10 {
//		return num
//	}
//	res := 0
//
//	for num > 0 {
//		res += num % 10
//		num /= 10
//	}
//
//	return addDigits(res)
//}

func addDigits(num int) int {
	if num == 0 {
		return num
	}
	if num%9 == 0 {
		return 9
	}
	return num % 9
}
