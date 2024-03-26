package find_all_numbers_disappeared_in_an_array

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFindDisappearedNumbers(t *testing.T) {
	nums := []int{4, 3, 2, 7, 8, 2, 3, 1}
	expected := []int{5, 6}

	assert.Equal(t, expected, findDisappearedNumbers(nums))
}

func TestFindDisappearedNumbers2(t *testing.T) {
	nums := []int{1, 1}
	expected := []int{2}

	assert.Equal(t, expected, findDisappearedNumbers(nums))
}
