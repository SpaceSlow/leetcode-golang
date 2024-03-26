package is_subsequence

func isSubsequence(s string, t string) bool {
	c := 0
	for i := range t {
		if c == len(s) {
			return true
		} else if s[c] == t[i] {
			c++
		}
	}

	return c == len(s)
}
