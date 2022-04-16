package problem

import (
	"math"
)

// https://leetcode.com/problems/maximum-subarray/
func MaxSubArray(nums []int) int {
	localMax := 0
	globalMax := math.MinInt
	// indices := make(map[int][]int)

	for i := 0; i < len(nums); i++ {
		if localMax+nums[i] >= nums[i] {
			// if i == 0 {
			// 	indices[i] = []int{0}
			// } else {
			// 	temp := append(indices[i-1], i)
			// 	indices[i] = temp
			// }
			localMax = localMax + nums[i]
		} else {
			// indices[i] = []int{i}
			localMax = nums[i]
		}
		if localMax >= globalMax {
			globalMax = localMax
		}
	}

	// for i := 0; i < len(indices); i++ {
	// 	fmt.Println(indices[i])
	// }

	return globalMax
}
