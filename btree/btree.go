package btree

// Node represent for a binary tree node.
type Node struct {
	Val   int
	left  *Node
	right *Node
}

func (n *Node) PreorderTraversal(f func(v int)) {
	if n == nil {
		return
	}

	n.left.PreorderTraversal(f)
	f(n.Val)
	n.right.PreorderTraversal(f)
}

func PreorderTraversal(root *Node) []int {
	arr := []int{root.Val}
	if root.left != nil {
		arr = append(arr, PreorderTraversal(root.left)...)
	}
	if root.right != nil {
		arr = append(arr, PreorderTraversal(root.right)...)
	}
	return arr
}
