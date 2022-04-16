package problem

import (
	"strconv"
)

// https://leetcode.com/problems/palindrome-number/
func IsPalindrome(x int) bool {
	xString := strconv.Itoa(x)
	lenX := len(xString)
	for i := 0; i < lenX/2; i++ {
		if xString[i] == xString[lenX-i-1] {
			continue
		}
		return false
	}
	return true
}
