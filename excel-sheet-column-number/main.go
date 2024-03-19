package excel_sheet_column_number

import "math"

//https://leetcode.com/problems/excel-sheet-column-number/

//A -> 1
//B -> 2
//C -> 3
//...
//Z -> 26
//AA -> 27
//AB -> 28
//...

func titleToNumber(columnTitle string) int {
	const zeroElement = 64
	const numberSystem = 26

	result := 0
	for i, char := range columnTitle {
		result += (int(char) - zeroElement) * int(math.Pow(numberSystem, float64(len(columnTitle)-i-1)))
	}
	return result
}
