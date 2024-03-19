package missing_number

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMissingNumber301(t *testing.T) {
	nums := []int{3, 0, 1}
	expected := 2

	assert.Equal(t, missingNumber(nums), expected)
}

func TestMissingNumber01(t *testing.T) {
	nums := []int{0, 1}
	expected := 2

	assert.Equal(t, missingNumber(nums), expected)
}

func TestMissingNumber0(t *testing.T) {
	nums := []int{0}
	expected := 1

	assert.Equal(t, missingNumber(nums), expected)
}

func TestMissingNumber1(t *testing.T) {
	nums := []int{1}
	expected := 0

	assert.Equal(t, missingNumber(nums), expected)
}

func TestMissingNumber964235701(t *testing.T) {
	nums := []int{9, 6, 4, 2, 3, 5, 7, 0, 1}
	expected := 8

	assert.Equal(t, missingNumber(nums), expected)
}
