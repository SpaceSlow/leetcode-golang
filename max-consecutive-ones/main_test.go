package max_consecutive_ones

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFindMaxConsecutiveOnesThreeOnesInEndAfterLocalMax(t *testing.T) {
	nums := []int{1, 1, 0, 1, 1, 1}
	expected := 3

	assert.Equal(t, expected, findMaxConsecutiveOnes(nums))
}

func TestFindMaxConsecutiveOnesTwoOnesInCenter(t *testing.T) {
	nums := []int{1, 0, 1, 1, 0, 1}
	expected := 2

	assert.Equal(t, expected, findMaxConsecutiveOnes(nums))
}
