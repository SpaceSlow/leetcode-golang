package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func inorderTraversal(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	var l []int
	l = append(l, inorderTraversal(root.Left)...)
	l = append(l, root.Val)
	l = append(l, inorderTraversal(root.Right)...)
	return l
}
