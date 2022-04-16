package problem

// https://leetcode.com/problems/implement-strstr/
func StrStr(haystack string, needle string) int {
	for i := 0; i < len(haystack); i++ {
		if haystack[i] == needle[0] && len(haystack)-i >= len(needle) {
			for j := 0; j < len(needle); j++ {
				if haystack[i+j] != needle[j] {
					break
				}
				if j+1 == len(needle) {
					return i
				}
			}
		}
	}
	return -1
}
