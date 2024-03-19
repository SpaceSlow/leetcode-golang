package pascals_triangle_ii

// https://leetcode.com/problems/pascals-triangle-ii/

//func getRow(rowIndex int) []int {
//	if rowIndex == 0 {
//		return []int{1}
//	}
//	tempRow := make([]int, rowIndex+1)
//	curRow := make([]int, rowIndex+1)
//
//	curRow[0], curRow[1], tempRow[0] = 1, 1, 1
//
//	for i := 2; i < rowIndex+1; i++ {
//		tempRow[0], tempRow[i] = 1, 1
//
//		for j := 1; j < (i+2)/2; j++ {
//			elem := curRow[j-1] + curRow[j]
//			tempRow[j], tempRow[i-j] = elem, elem
//		}
//		tempRow, curRow = curRow, tempRow
//	}
//
//	return curRow
//}

// https://en.wikipedia.org/wiki/Pascal%27s_triangle#Calculating_a_row_or_diagonal_by_itself
func getRow(rowIndex int) []int {
	row := make([]int, rowIndex+1)

	row[0], row[rowIndex] = 1, 1

	for i := 1; i < rowIndex+1; i++ {
		row[i] = row[i-1] * (rowIndex - i + 1) / i
	}

	return row
}
