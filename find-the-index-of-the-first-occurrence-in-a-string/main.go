package find_the_index_of_the_first_occurrence_in_a_string

//https://leetcode.com/problems/find-the-index-of-the-first-occurrence-in-a-string/

//func strStr(haystack string, needle string) int {
//	for i := 0; i < len(haystack)-len(needle)+1; i++ {
//		for j := 0; j < len(needle); j++ {
//			if haystack[i] != needle[j] {
//				i -= j
//				break
//			}
//			if j == len(needle)-1 {
//				return i - j
//			}
//			i++
//		}
//	}
//	return -1
//}

func strStr(haystack string, needle string) int {
	if len(haystack) < len(needle) {
		return -1
	}
	for i := 0; i < len(haystack)-len(needle)+1; i++ {
		if haystack[i] == needle[0] && haystack[i:i+len(needle)] == needle {
			return i
		}
	}
	return -1
}
