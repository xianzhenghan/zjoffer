package zj054__二叉搜索树的第k大节点

//https://colobu.com/2020/07/15/implement-bst-in-Go/
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func kthLargest(root *TreeNode, k int) int {
	// 中序遍历 记录次数
	if root == nil {
		return -1
	}
	nums := make([]int, 0)
	midOrde(root, &nums)
	return nums[len(nums)-k]
}

func midOrde(node *TreeNode, nums *[]int) {
	if node.Left != nil {
		midOrde(node.Left, nums)
	}
	*nums = append(*nums, node.Val)
	if node.Right != nil {
		midOrde(node.Right, nums)
	}
}

func kthLargest2(root *TreeNode, k int) int {
	// 中序遍历 记录次数
	if root == nil {
		return -1
	}
	var dsf func(*TreeNode)
	res := -1
	dsf = func(node *TreeNode) {
		if node == nil {
			return
		}
		dsf(node.Right)
		k--
		if k == 0 {
			res = node.Val
		}
		dsf(node.Left)
	}
	dsf(root)
	return res
}
