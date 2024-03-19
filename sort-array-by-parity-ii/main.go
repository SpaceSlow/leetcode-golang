package sort_array_by_parity_ii

// https://leetcode.com/problems/sort-array-by-parity-ii/

//func sortArrayByParityII(nums []int) []int {
//	result := make([]int, len(nums))
//	evenI, oddsI := 0, 1
//
//	for i := 0; i < len(nums); i++ {
//		if nums[i]%2 == 0 {
//			result[evenI] = nums[i]
//			evenI += 2
//		} else {
//			result[oddsI] = nums[i]
//			oddsI += 2
//		}
//	}
//
//	return nums
//}

func sortArrayByParityII(nums []int) []int {

	for even, odd := 0, 1; even < len(nums) && odd < len(nums); {
		if nums[even]%2 != 0 && nums[odd]%2 == 0 {
			nums[even], nums[odd] = nums[odd], nums[even]
		}
		if nums[even]&1 == 0 {
			even += 2
		}
		if nums[odd]&1 == 1 {
			odd += 2
		}
	}

	return nums
}
