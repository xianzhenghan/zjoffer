package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func mirrorTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	return mirrorSubSymmetric(root.Left, root.Right)
}

func mirrorSubSymmetric(left *TreeNode, right *TreeNode) *TreeNode {

}
