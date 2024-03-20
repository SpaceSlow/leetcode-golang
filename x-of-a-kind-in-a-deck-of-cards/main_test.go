package x_of_a_kind_in_a_deck_of_cards

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestHasGroupsSizeXHasGroup2Size(t *testing.T) {
	deck := []int{1, 2, 3, 4, 4, 3, 2, 1}
	expected := true

	assert.Equal(t, expected, hasGroupsSizeX(deck))
}

func TestHasGroupsSizeXNoPossiblePartition(t *testing.T) {
	deck := []int{1, 1, 1, 2, 2, 2, 3, 3}
	expected := false

	assert.Equal(t, expected, hasGroupsSizeX(deck))
}
