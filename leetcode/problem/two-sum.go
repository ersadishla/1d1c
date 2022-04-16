package problem

// https://leetcode.com/problems/two-sum/
func TwoSum(nums []int, target int) []int {
	temp := map[int]int{}
	res := []int{0, 0}
	for i, num := range nums {
		temp[num] = i
	}

	for i, num := range nums {
		res[0] = i
		if temp[target-num] != 0 && temp[target-num] != i {
			res[1] = temp[target-num]
			break
		}
	}

	return res
}
