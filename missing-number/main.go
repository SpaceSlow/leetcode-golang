package missing_number

// https://leetcode.com/problems/missing-number/description/

func missingNumber(nums []int) int {
	missed := len(nums)

	for i, num := range nums {
		missed ^= i ^ num
	}
	return missed
}
