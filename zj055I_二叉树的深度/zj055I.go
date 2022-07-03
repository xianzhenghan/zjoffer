package zj055I_二叉树的深度

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxDepth(root *TreeNode) int {
	return dfs(root)
}

func dfs(node *TreeNode) int {
	if node == nil {
		return 0
	}

	return int(math.Max(float64(dfs(node.Left)), float64(dfs(node.Right)))) + 1
}
