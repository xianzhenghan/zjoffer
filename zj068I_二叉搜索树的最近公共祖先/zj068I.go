package zj068I_二叉搜索树的最近公共祖先

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func lowestCommonAncestor(root *TreeNode, p, q *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	if p.Val > q.Val {
		return lowestCommonAncestor(root, q, p)
	}
	if p.Val <= root.Val && root.Val <= q.Val {
		return root
	}
	if root.Val < p.Val {
		return lowestCommonAncestor(root.Left, p, q)
	} else {
		return lowestCommonAncestor(root.Right, p, q)
	}
}

func lowestCommonAncestor2(root *TreeNode, p, q *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	if p.Val > q.Val {
		return lowestCommonAncestor(root, q, p)
	}
	if p.Val <= root.Val && root.Val <= q.Val {
		return root
	}
	if root.Val < p.Val {
		return lowestCommonAncestor(root.Left, p, q)
	} else {
		return lowestCommonAncestor(root.Right, p, q)
	}
}
