package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//func inorderTraversal(root *TreeNode) []int {
//	if root == nil {
//		return []int{}
//	}
//	var l []int
//	l = append(l, inorderTraversal(root.Left)...)
//	l = append(l, root.Val)
//	l = append(l, inorderTraversal(root.Right)...)
//	return l
//}

func inorderTraversal(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}

	deque := make([]*TreeNode, 0)

	nodes := []int{}
	for root != nil || len(deque) != 0 {
		for root != nil {
			deque = append(deque, root)
			root = root.Left
		}
		lastNode := len(deque) - 1
		root = deque[lastNode]
		deque = deque[:lastNode]
		nodes = append(nodes, root.Val)
		root = root.Right
	}

	return nodes
}
