package find_all_numbers_disappeared_in_an_array

// https://leetcode.com/problems/find-all-numbers-disappeared-in-an-array/

//func findDisappearedNumbers(nums []int) []int {
//	visited := make(map[int]bool, len(nums))
//
//	for i := 0; i < len(nums); i++ {
//		visited[nums[i]] = true
//	}
//
//	result := make([]int, 0)
//	for i := 0; i < len(nums); i++ {
//		if _, ok := visited[i+1]; !ok {
//			result = append(result, i+1)
//		}
//	}
//
//	return result
//}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func findDisappearedNumbers(nums []int) []int {
	for _, num := range nums {
		i := abs(num) - 1
		if nums[i] > 0 {
			nums[i] = -nums[i]
		}
	}

	result := make([]int, 0)
	for i := 0; i < len(nums); i++ {
		if nums[i] > 0 {
			result = append(result, i+1)
		}
	}

	return result
}
