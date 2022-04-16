package problem

// https://leetcode.com/problems/longest-common-prefix/
func LongestCommonPrefix(strs []string) string {
	firstStr := string(strs[0])
	firstStrLen := len(strs[0])
	res := ""
	for i := 0; i < firstStrLen; i++ {
		for j := 1; j < len(strs); j++ {
			if len(res) == len(strs[j]) || string(firstStr[i]) != string(strs[j][i]) {
				return res
			}
		}
		res += string(strs[0][i])
	}
	return res
}
