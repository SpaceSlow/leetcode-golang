package excel_sheet_column_title

func convertToTitle(columnNumber int) string {
	result := ""

	for ; columnNumber > 0; columnNumber /= 26 {
		columnNumber--
		result = string(rune(columnNumber%26+65)) + result
	}

	return result
}
