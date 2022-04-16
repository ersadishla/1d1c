package main

import (
	"1d1c/leetcode/problem"
	"fmt"
)

func main() {
	fmt.Println("=== MAIN PACKAGE ===")
	// sandbox()
	// twoSum()
	// isPalindrome()
	// longestCommonPrefix()
	// isValidParentheses()
	// mergeTwoLists()
	// removeDuplicates()
	// removeElement()
	// strStr()
	// searchInsert()
	// maxSubArray()
	// lengthOfLastWord()
}

func twoSum() {
	nums := []int{1, 2, 3}

	fmt.Println(problem.TwoSum(nums, 3))
}

func isPalindrome() {
	x := 11233211

	fmt.Println(problem.IsPalindrome(x))
}

func longestCommonPrefix() {
	strs := []string{
		"flower",
		"flow",
		"flight",
	}

	fmt.Println(problem.LongestCommonPrefix(strs))
}

func isValidParentheses() {
	s := "}{}"

	fmt.Println(problem.IsValidParentheses(s))
}

func mergeTwoLists() {
	l0 := new(problem.ListNode)
	l0.Val = 0
	l0.Next = nil
	l1 := new(problem.ListNode)
	l1.Val = 1
	l1.Next = nil
	l2 := new(problem.ListNode)
	l2.Val = 2
	l2.Next = nil
	l3 := new(problem.ListNode)
	l3.Val = 3
	l3.Next = nil
	l4 := new(problem.ListNode)
	l4.Val = 4
	l4.Next = nil
	l5 := new(problem.ListNode)
	l5.Val = 5
	l5.Next = nil
	l6 := new(problem.ListNode)
	l6.Val = 6
	l6.Next = nil

	l1.Next = l2
	l2.Next = l4

	head := problem.MergeTwoLists(l5, l1)
	for {
		fmt.Println(head.Val)
		head = head.Next

		if head == nil {
			break
		}
	}
}

func removeDuplicates() {
	nums := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}

	k := problem.RemoveDuplicates(nums)

	for i := 0; i < k; i++ {
		fmt.Println(nums[i])
	}
}

func removeElement() {
	nums := []int{0, 1, 2, 2, 3, 0, 4, 2, 5, 2}

	k := problem.RemoveElement(nums, 2)

	for i := 0; i < k; i++ {
		fmt.Println(nums[i])
	}
}

func strStr() {
	fmt.Println(problem.StrStr("mississippi", "issip"))
}

func searchInsert() {
	nums := []int{1, 2, 4, 6, 8, 10}

	fmt.Println(problem.SearchInsert(nums, 11))
}

func maxSubArray() {
	nums := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}

	fmt.Println(problem.MaxSubArray(nums))
}

func lengthOfLastWord() {
	s := "   fly me   to   the moon  "

	fmt.Println(problem.LengthOfLastWord(s))
}
