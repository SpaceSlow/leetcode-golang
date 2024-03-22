package number_of_beautiful_pairs

// https://leetcode.com/problems/number-of-beautiful-pairs/

func countBeautifulPairs(nums []int) int {
	counter := 0

	for i := 0; i < len(nums)-1; i++ {
		firstDigit := getFirstDigit(nums[i])
		for j := i + 1; j < len(nums); j++ {
			if gcd(firstDigit, getLastDigit(nums[j])) == 1 {
				counter++
			}
		}
	}

	return counter
}

func getFirstDigit(n int) int {
	for ; n >= 10; n /= 10 {
	}
	return n
}

func getLastDigit(n int) int {
	return n % 10
}

func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}
