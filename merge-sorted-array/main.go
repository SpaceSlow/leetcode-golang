package merge_sorted_array

// https://leetcode.com/problems/merge-sorted-array/

func merge(nums1 []int, m int, nums2 []int, n int) {
	if n == 0 {
		return
	}
	i, j := m, n
	for k := m + n - 1; i > 0 || j > 0; k-- {
		if j > 0 && (i == 0 || nums2[j-1] > nums1[i-1]) {
			j--
			nums1[k] = nums2[j]
		} else {
			i--
			nums1[k] = nums1[i]
		}
	}
}
