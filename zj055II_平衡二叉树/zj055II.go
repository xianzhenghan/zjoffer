package zj055II_二叉树的深度

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return dsf(root)

}

func dsf(node *TreeNode) bool {
	if node == nil {
		return true
	}
	if math.Abs(float64(deep(node.Left)-deep(node.Right))) > 1 {
		return false
	}
	return dsf(node.Left) && dsf(node.Right)
}

func deep(node *TreeNode) int {
	if node == nil {
		return 0
	}
	return int(math.Max(float64(deep(node.Left)), float64(deep(node.Right)))) + 1
}
