package two_sum_iv_input_is_a_bst

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func findTarget(root *TreeNode, k int) bool {
	visited := make(map[int]bool)

	var dfs func(root *TreeNode) bool
	dfs = func(root *TreeNode) bool {
		if root == nil {
			return false
		} else if visited[k-root.Val] {
			return true
		} else {
			visited[root.Val] = true
		}

		return dfs(root.Left) || dfs(root.Right)
	}

	return dfs(root)
}
