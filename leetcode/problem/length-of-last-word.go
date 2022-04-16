package problem

// https://leetcode.com/problems/length-of-last-word/
func LengthOfLastWord(s string) int {
	c := 0
	detected := false
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] != 32 {
			c++
			detected = true
		}
		if detected && s[i] == 32 {
			break
		}
	}

	return c
}
