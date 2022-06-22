package zj032I_从上到下打印二叉树

import "container/list"

//*
//* Definition for a binary tree node.
//* type TreeNode struct {
//*     Val int
//*     Left *TreeNode
//*     Right *TreeNode
//* }

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrder(root *TreeNode) []int {
	var nums = make([]int, 0)
	if root == nil {
		return nums
	}
	queue := list.New()
	queue.PushBack(root)
	for queue.Len() != 0 {
		node := queue.Front().Value.(*TreeNode)
		queue.Remove(queue.Front())
		nums = append(nums, node.Val)
		if node.Left != nil {
			queue.PushBack(node.Left)
		}
		if node.Right != nil {
			queue.PushBack(node.Right)
		}
	}
	return nums
}
