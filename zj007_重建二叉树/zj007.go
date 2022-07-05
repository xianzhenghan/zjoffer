package zj007_重建二叉树

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}
	return dsf(preorder, inorder)
}

func dsf(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	} else {
		val_preorder := preorder[0]
		root := &TreeNode{
			Val:   val_preorder,
			Left:  nil,
			Right: nil,
		}
		idx_inorder := getIndex(val_preorder, inorder)
		left := dsf(preorder[1:idx_inorder+1], inorder[0:idx_inorder])
		right := dsf(preorder[idx_inorder+1:], inorder[idx_inorder+1:])
		root.Left, root.Right = left, right
		return root
	}
}

// k的位置，k 之前的长度，k之后的长度
func getIndex(val int, nums []int) int {
	for k, v := range nums {
		if val == v {
			return k
		}
	}
	return -1
}
