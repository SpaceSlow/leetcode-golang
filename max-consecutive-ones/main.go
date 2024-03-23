package max_consecutive_ones

// https://leetcode.com/problems/max-consecutive-ones/description/

func findMaxConsecutiveOnes(nums []int) int {
	maxConsecutiveOnes := 0

	counterConsecutiveOnes := 0
	for _, num := range nums {
		if num == 1 {
			counterConsecutiveOnes++
			continue
		}
		maxConsecutiveOnes = max(counterConsecutiveOnes, maxConsecutiveOnes)
		counterConsecutiveOnes = 0
	}

	return max(counterConsecutiveOnes, maxConsecutiveOnes)
}
