package determine_the_winner_of_a_bowling_game

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIsWinnerWithStrike(t *testing.T) {
	player1 := []int{4, 10, 7, 9}
	player2 := []int{6, 5, 2, 3}
	exptected := 1

	assert.Equal(t, exptected, isWinner(player1, player2))
}

func TestIsWinnerDraw(t *testing.T) {
	player1 := []int{2, 3}
	player2 := []int{4, 1}
	exptected := 0

	assert.Equal(t, exptected, isWinner(player1, player2))
}
