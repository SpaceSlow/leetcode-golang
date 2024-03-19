package nim_game

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCanWinNim4(t *testing.T) {
	rocksNum := 4
	expected := false

	assert.Equal(t, expected, canWinNim(rocksNum))
}

func TestCanWinNim1(t *testing.T) {
	rocksNum := 1
	expected := true

	assert.Equal(t, expected, canWinNim(rocksNum))
}

func TestCanWinNim2(t *testing.T) {
	rocksNum := 2
	expected := true

	assert.Equal(t, expected, canWinNim(rocksNum))
}

func TestCanWinNim8(t *testing.T) {
	rocksNum := 8
	expected := false

	assert.Equal(t, expected, canWinNim(rocksNum))
}

func TestCanWinNim6(t *testing.T) {
	rocksNum := 6
	expected := true

	assert.Equal(t, expected, canWinNim(rocksNum))
}

func TestCanWinNim5(t *testing.T) {
	rocksNum := 5
	expected := true

	assert.Equal(t, expected, canWinNim(rocksNum))
}
