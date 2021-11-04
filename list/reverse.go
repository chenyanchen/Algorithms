// Copyright 2020 Singularity, Inc. All rights reserved.

package list

// Node represent a node of the singly-linked list.
type Node struct {
	Val  int
	Next *Node
}

func Reverse(head *Node) *Node {
	var (
		cur      = head
		tmp, pre *Node
	)

	for cur != nil {
		tmp = cur.Next
		cur.Next = pre
		pre = cur
		cur = tmp
	}

	return pre
}
