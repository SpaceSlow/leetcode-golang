package magic_squares_in_grid

// https://leetcode.com/problems/magic-squares-in-grid/

func numMagicSquaresInside(grid [][]int) int {
	countMagicSquares := 0
	for i := 0; i <= len(grid)-3; i++ {
		for j := 0; j <= len(grid[0])-3; j++ {
			if checkMagicSquare(grid, i, j) {
				countMagicSquares++
			}
		}
	}

	return countMagicSquares
}

func checkMagicSquare(grid [][]int, column, row int) bool {
	visited := make(map[int]bool)
	for i := column; i < column+3; i++ {
		for j := row; j < row+3; j++ {
			if visited[grid[i][j]] || grid[i][j] > 9 || grid[i][j] == 0 {
				return false
			}
			visited[grid[i][j]] = true
		}
	}

	for i := column; i < column+3; i++ {
		if grid[column][row]+grid[column][row+1]+grid[column][row+2] != 15 {
			return false
		}
	}

	for j := row; j < row+3; j++ {
		if grid[column][row]+grid[column+1][row]+grid[column+2][row] != 15 {
			return false
		}
	}

	if grid[column][row]+grid[column+1][row+1]+grid[column+2][row+2] != 15 {
		return false
	}
	if grid[column+2][row]+grid[column+1][row+1]+grid[column][row+2] != 15 {
		return false
	}

	return true
}
