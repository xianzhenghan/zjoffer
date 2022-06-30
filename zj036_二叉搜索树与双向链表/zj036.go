package zj036_二叉搜索树与双向链表

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func treeToDoublyList(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	head := &TreeNode{}
	pre := &TreeNode{}
	dsf(root, head, pre)
	head.Left = pre
	pre.Right = head
	return head
}

func dsf(node *TreeNode, head, pre *TreeNode) {
	if node == nil {
		return
	}
	dsf(node.Left, head, pre)
	if head == nil {
		head = node
		pre = node
	} else {
		node.Left = pre
		pre.Right = node
		pre = node
	}
	dsf(node.Right, head, pre)
}
