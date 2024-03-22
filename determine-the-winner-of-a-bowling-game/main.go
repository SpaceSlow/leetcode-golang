package determine_the_winner_of_a_bowling_game

// https://leetcode.com/problems/determine-the-winner-of-a-bowling-game/

func isWinner(player1 []int, player2 []int) int {
	scoreFirstPlayer, scoreSecondPlayer := 0, 0

	for i := 0; i < len(player1); i++ {
		scoreFirstPlayer += getPointOnTurn(player1, i)
		scoreSecondPlayer += getPointOnTurn(player2, i)
	}

	if scoreFirstPlayer > scoreSecondPlayer {
		return 1
	} else if scoreFirstPlayer < scoreSecondPlayer {
		return 2
	}
	return 0
}

func getPointOnTurn(points []int, turn int) int {
	if turn > 0 && points[turn-1] == 10 || turn > 1 && points[turn-2] == 10 {
		return points[turn] * 2
	}
	return points[turn]
}
