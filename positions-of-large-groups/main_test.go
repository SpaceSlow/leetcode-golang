package positions_of_large_groups

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLargeGroupPositionsOneGroup(t *testing.T) {
	s := "abbxxxxzzy"
	expected := [][]int{{3, 6}}

	assert.Equal(t, expected, largeGroupPositions(s))
}

func TestLargeGroupPositionsNoOneGroup(t *testing.T) {
	s := "abc"
	expected := [][]int{}

	assert.Equal(t, expected, largeGroupPositions(s))
}

func TestLargeGroupPositionsSeveralGroups(t *testing.T) {
	s := "abcdddeeeeaabbbcd"
	expected := [][]int{{3, 5}, {6, 9}, {12, 14}}

	assert.Equal(t, expected, largeGroupPositions(s))
}
