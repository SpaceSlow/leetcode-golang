package move_zeroes

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMoveZeroes(t *testing.T) {
	nums := []int{0, 0, 1}
	expected := []int{1, 0, 0}

	moveZeroes(nums)
	assert.Equal(t, expected, nums)
}
