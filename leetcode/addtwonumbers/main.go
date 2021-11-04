package addtwonumbers

// Definition for singly-linked list.
type Node struct {
	Val  int
	Next *Node
}

func addTwoNumbers(l1 *Node, l2 *Node) *Node {
	root := &Node{}
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
			node.Next = &Node{Val: node.Val / 10}
			node.Val %= 10
		} else if node1 == nil && node2 == nil {
			break
		} else {
			node.Next = &Node{}
		}
		node = node.Next
	}
	return root
}
