package search_insert_position

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSearchInsert(t *testing.T) {
	nums := []int{1, 3, 5, 6}
	target := 2
	expected := 1

	assert.Equal(t, expected, searchInsert(nums, target))
}
