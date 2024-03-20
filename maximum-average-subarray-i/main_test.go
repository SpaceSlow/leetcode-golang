package maximum_average_subarray_i

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFindMaxAverageSeveralNums(t *testing.T) {
	nums := []int{1, 12, -5, -6, 50, 3}
	k := 4
	expected := 12.75000

	assert.Equal(t, expected, findMaxAverage(nums, k))
}

func TestFindMaxAverageOneNum(t *testing.T) {
	nums := []int{5}
	k := 1
	expected := 5.00000

	assert.Equal(t, expected, findMaxAverage(nums, k))
}

func TestFindMaxAverageNegativeNum(t *testing.T) {
	nums := []int{-1}
	k := 1
	expected := -1.00000

	assert.Equal(t, expected, findMaxAverage(nums, k))
}
