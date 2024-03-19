package positions_of_large_groups

// https://leetcode.com/problems/positions-of-large-groups/description/

func largeGroupPositions(s string) [][]int {
	groups := make([][]int, 0)

	if len(s) < 3 {
		return groups
	}

	counter := 1
	for i := 1; i < len(s); i++ {
		if s[i] == s[i-1] {
			counter++
		} else if counter >= 3 {
			groups = append(groups, []int{i - counter, i - 1})
			counter = 1
		} else {
			counter = 1
		}
	}

	if counter >= 3 {
		groups = append(groups, []int{len(s) - counter, len(s) - 1})
	}

	return groups
}
