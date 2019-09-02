package leetcode

//package leetcode

// You are given two non-empty linked lists representing two non-negative integers.
// The digits are stored in reverse order and each of their nodes contain a single digit.
// Add the two numbers and return it as a linked list.
//
// You may assume the two numbers do not contain any leading zero, except the number 0 itself.
//
// Example:
//
//
//Input: (2 -> 4 -> 3) + (5 -> 6 -> 4)
//Output: 7 -> 0 -> 8
//Explanation: 342 + 465 = 807.
//
//

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	return addTwoNumbersRec(l1, l2, false)
}

func addTwoNumbersRec(l1 *ListNode, l2 *ListNode, prevUp bool) *ListNode {
	if l1 == nil && l2 == nil {
		if prevUp {
			return &ListNode{Val: 1, Next: nil}
		}
		return nil
	}

	e1, e2 := 0, 0
	var next1, next2 *ListNode

	if l1 != nil {
		e1 = l1.Val
		next1 = l1.Next
	}
	if l2 != nil {
		e2 = l2.Val
		next2 = l2.Next
	}

	sum := e1 + e2

	if prevUp {
		sum += 1
	}

	up := false
	if sum > 9 {
		sum -= 10
		up = true
	}

	return &ListNode{
		Val:  sum,
		Next: addTwoNumbersRec(next1, next2, up),
	}
}
