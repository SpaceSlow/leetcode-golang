package summary_ranges

import "strconv"

// https://leetcode.com/problems/summary-ranges/

func summaryRanges(nums []int) []string {
	if len(nums) == 0 {
		return []string{}
	}

	start := nums[0]
	ranges := make([]string, 0)
	for i := 1; i < len(nums); i++ {
		if nums[i-1]+1 == nums[i] {
			continue
		} else if start == nums[i-1] {
			ranges = append(ranges, strconv.Itoa(start))
		} else {
			ranges = append(ranges, strconv.Itoa(start)+"->"+strconv.Itoa(nums[i-1]))
		}
		start = nums[i]
	}

	if start == nums[len(nums)-1] {
		ranges = append(ranges, strconv.Itoa(start))
	} else {
		ranges = append(ranges, strconv.Itoa(start)+"->"+strconv.Itoa(nums[len(nums)-1]))
	}

	return ranges
}
