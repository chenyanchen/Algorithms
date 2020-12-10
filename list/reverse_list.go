// Copyright 2020 Singularity, Inc. All rights reserved.

package list

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	var (
		cur      = head
		tmp, pre *ListNode
	)

	for cur != nil {
		tmp = cur.Next
		cur.Next = pre
		pre = cur
		cur = tmp
	}

	return pre
}
