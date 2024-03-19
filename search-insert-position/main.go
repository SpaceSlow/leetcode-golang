package search_insert_position

// https://leetcode.com/problems/search-insert-position/

//func searchInsert(nums []int, target int) int {
//	if target == nums[len(nums)/2] {
//		return len(nums) / 2
//	} else if len(nums) == 1 {
//		if target > nums[0] {
//			return 1
//		} else {
//			return 0
//		}
//	} else if target > nums[len(nums)/2] {
//		return searchInsert(nums[len(nums)/2:], target) + len(nums)/2
//	} else if target < nums[len(nums)/2] {
//		return searchInsert(nums[:len(nums)/2], target)
//	}
//	return len(nums)
//}

func searchInsert(nums []int, target int) int {
	i, j := 0, len(nums)
	for i < j {
		middle := (j-i)/2 + i

		if nums[middle] == target {
			return middle
		} else if target > nums[middle] {
			i = middle + 1
		} else {
			j = middle
		}
	}
	return i
}
