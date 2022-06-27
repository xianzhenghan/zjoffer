package zj028_对称二叉树

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSymmetric(root *TreeNode) bool {
	return IsMirrorTree(root, root)
}

func IsMirrorTree(l *TreeNode, r *TreeNode) bool {
	if l == nil && r == nil {
		return true
	}
	if l != nil || r != nil {
		return false
	}
	return l.Val == r.Val && IsMirrorTree(l.Left, r.Right) && IsMirrorTree(l.Right, r.Left)
}
