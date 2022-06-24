package zj032II_从上到下打印二叉树II

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

func levelOrder(root *TreeNode) [][]int {
	var nums = make([][]int, 0)
	if root == nil {
		return [][]int{}
	}
	queue := list.New()
	queue.PushBack(root)
	for queue.Len() != 0 {
		//输出这层的数组，把这一层的数据如队列
		vals, nextLevelqueue := levelNodes(queue)
		nums = append(nums, vals)
		queue = nextLevelqueue
	}
	return nums
}
func levelNodes(q *list.List) ([]int, *list.List) {
	queue := list.New()
	vals := make([]int, 0)
	for q.Len() != 0 {
		node := q.Front().Value.(*TreeNode)
		vals = append(vals, node.Val)
		q.Remove(q.Front())
		if node.Left != nil {
			queue.PushBack(node.Left)
		}
		if node.Right != nil {
			queue.PushBack(node.Right)
		}
	}
	return vals, queue
}
