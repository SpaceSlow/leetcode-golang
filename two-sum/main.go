package main

import "fmt"

func twoSum(nums []int, target int) []int {
	res := make([]int, 2)
	for i := 0; i < len(nums)-1; i++ {
		for j := i + 1; j < len(nums)-1; j++ {
			if nums[i]+nums[j] == target {
				res[0], res[1] = nums[i], nums[j]
			}
		}
	}
	return res
}

func main() {
	nums := []int{2, 7, 11, 15}
	target := 9

	fmt.Println(twoSum(nums, target))
}
