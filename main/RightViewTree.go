package main

import "fmt"

var rightViewTreeSequence [][]int
var rightView []int

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func rightSideView(root *TreeNode) []int {
	rightViewTreeSequence = make([][]int, 0)
	rightView = make([]int, 0)
	levelOrderExec2(root, 0)
	return rightView
}
func levelOrderExec2(root *TreeNode, lay int) {
	// lay 表示当前树节点所在的层 (从0开始)
	if root == nil {
		return
	}
	if len(rightViewTreeSequence) == lay {
		rightViewTreeSequence = append(rightViewTreeSequence, []int{})
		rightView = append(rightView, 0)
	}
	rightViewTreeSequence[lay] = append(rightViewTreeSequence[lay], root.Val)
	rightView[lay] = root.Val
	levelOrderExec2(root.Left, lay+1)
	levelOrderExec2(root.Right, lay+1)
}

func main() {
	fmt.Println(rightSideView(&TreeNode{1, &TreeNode{2, &TreeNode{3, nil, nil}, &TreeNode{4, nil, nil}}, nil}))
}
