// Copyright 2020 Singularity, Inc. All rights reserved.

// Reversion:
//   Init: Jon Snow    2020/4/16 22:42

package addtwonumbers

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	root := &ListNode{}
	node, node1, node2 := root, l1, l2
	for {
		var v1, v2 int
		if node1 != nil {
			v1 = node1.Val
			node1 = node1.Next
		}
		if node2 != nil {
			v2 = node2.Val
			node2 = node2.Next
		}
		node.Val += v1 + v2
		if node.Val >= 10 {
			node.Next = &ListNode{Val: node.Val / 10}
			node.Val = node.Val % 10
		} else if node1 == nil && node2 == nil {
			break
		} else {
			node.Next = &ListNode{}
		}
		node = node.Next
	}
	return root
}
