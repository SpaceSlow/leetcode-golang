package sort_array_by_parity_ii

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSortArrayByParittyIIStartEven(t *testing.T) {
	nums := []int{4, 2, 5, 7}
	expected := []int{4, 5, 2, 7}

	assert.Equal(t, expected, sortArrayByParityII(nums))
}

func TestSortArrayByParittyIIStartOdd(t *testing.T) {
	nums := []int{5, 7, 4, 2}
	expected := []int{2, 7, 4, 5}

	assert.Equal(t, expected, sortArrayByParityII(nums))
}

func TestSortArrayByParittyIINeededReshuffle(t *testing.T) {
	nums := []int{4, 1, 2, 1}
	expected := []int{4, 1, 2, 1}

	assert.Equal(t, expected, sortArrayByParityII(nums))
}

func TestSortArrayByParittyIIWithoutReshuffle(t *testing.T) {
	nums := []int{2, 3}
	expected := []int{2, 3}

	assert.Equal(t, expected, sortArrayByParityII(nums))
}
