package main

import "fmt"

type TreeNode struct {
	Left  *TreeNode
	Right *TreeNode
	Val   int
}

func buildTree(preorder []int, inorder []int) *TreeNode {
	return PrePostOrderBuildTree(preorder, inorder)
}

func PrePostOrderBuildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}
	if len(preorder) == 1 {
		return &TreeNode{nil, nil, preorder[0]}
	}
	index, value := 0, preorder[0]
	for i := 0; i < len(inorder); i++ {
		if value == inorder[i] {
			index = i
		}
	}
	return &TreeNode{
		Left:  PrePostOrderBuildTree(preorder[1:index+1], inorder[:index]),
		Right: PrePostOrderBuildTree(preorder[index+1:], inorder[index+1:]),
		Val:   value,
	}
}

func main() {
	/*	left := &TreeNode{Left: &TreeNode{nil, nil, 3}, Right: &TreeNode{nil, nil, 4}, Val: 2}
		right := &TreeNode{Left: &TreeNode{nil, nil, 6}, Right: nil, Val: 5}
		root := &TreeNode{left, right, 1}
		fmt.Println(levelOrder(root))*/

	preorder := []int{1, 2, 3, 4, 5, 6}
	inorder := []int{2, 3, 4, 1, 6, 5}
	root := (PrePostOrderBuildTree(preorder, inorder))
	fmt.Println(levelOrder(root))
}

var levelOrderSequence [][]int

func levelOrder(root *TreeNode) [][]int {
	levelOrderSequence = make([][]int, 0)
	levelOrderExec(root, 0)
	return levelOrderSequence
}
func levelOrderExec(root *TreeNode, lay int) {
	// lay 表示当前树节点所在的层 (从0开始)
	if root == nil {
		return
	}
	if len(levelOrderSequence) == lay {
		levelOrderSequence = append(levelOrderSequence, []int{})
	}
	levelOrderSequence[lay] = append(levelOrderSequence[lay], root.Val)
	levelOrderExec(root.Left, lay+1)
	levelOrderExec(root.Right, lay+1)
}
