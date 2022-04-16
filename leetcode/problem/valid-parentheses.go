package problem

// https://leetcode.com/problems/valid-parentheses/
func IsValidParentheses(s string) bool {
	var stemp []string
	for _, c := range s {
		if string(c) == "(" || string(c) == "[" || string(c) == "{" {
			stemp = append(stemp, string(c))
		}
		if string(c) == ")" || string(c) == "]" || string(c) == "}" {
			if len(stemp) == 0 {
				return false
			}
			bracket := stemp[len(stemp)-1]
			stemp = stemp[:len(stemp)-1]
			if bracket == "" {
				return false
			}
			if bracket == "(" && string(c) != ")" {
				return false
			}
			if bracket == "[" && string(c) != "]" {
				return false
			}
			if bracket == "{" && string(c) != "}" {
				return false
			}
		}
	}
	return len(stemp) == 0
}
