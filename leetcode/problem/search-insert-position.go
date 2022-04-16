package problem

// https://leetcode.com/problems/search-insert-position/
func SearchInsert(nums []int, target int) int {
	divider := len(nums) / 2

	if nums[divider] == target {
		return divider
	}
	if target > nums[len(nums)-1] {
		return len(nums)
	}
	if target < nums[0] {
		return 0
	}
	if target < nums[divider] {
		for i := 0; i < divider; i++ {
			if nums[i] >= target {
				return i
			}
			if nums[i+1] >= target {
				return i + 1
			}
		}
	} else {
		for i := divider + 1; i < len(nums); i++ {
			if nums[i] >= target {
				return i
			}
			if nums[i+1] >= target {
				return i + 1
			}
		}
	}

	return 0
}
