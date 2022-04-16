package problem

import "strconv"

// https://leetcode.com/problems/plus-one/
func PlusOne(digits []int) []int {
	var strFormat string
	var intFormat int
	res := make([]int, 0)

	for i := 0; i < len(digits); i++ {
		strFormat += strconv.Itoa(digits[i])
	}

	intFormat, _ = strconv.Atoi(strFormat)
	intFormat++
	strFormat = strconv.Itoa(intFormat)

	for i := 0; i < len(strFormat); i++ {
		intTemp, _ := strconv.Atoi(string(strFormat[i]))
		res = append(res, intTemp)
	}

	return res
}
