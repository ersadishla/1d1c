package problem

/**
 * Definition for singly-linked list.
 */
type ListNode struct {
	Val  int
	Next *ListNode
}

// https://leetcode.com/problems/merge-two-sorted-lists/
func MergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	var head = new(ListNode)
	var temp = new(ListNode)
	if list1 == nil && list2 == nil {
		return nil
	}
	if list1 == nil {
		return list2
	}
	if list2 == nil {
		return list1
	}
	if list1 != nil && list1.Val <= list2.Val {
		temp = list1
		head = temp
		list1 = list1.Next
	} else {
		temp = list2
		head = temp
		list2 = list2.Next
	}
	for {
		if list2 == nil || (list1 != nil && list1.Val <= list2.Val) {
			temp.Next = list1
			temp = list1
			list1 = list1.Next
		} else if list2 != nil {
			temp.Next = list2
			temp = list2
			list2 = list2.Next
		}

		if list1 == nil && list2 == nil {
			break
		}
	}

	return head
}
