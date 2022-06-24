package zj032III_从上到下打印二叉树III

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
	level := 0
	for queue.Len() != 0 {
		//输出这层的数组，把这一层的数据如队列
		vals, nextLevelqueue := levelNodes(queue, level)
		nums = append(nums, vals)
		queue = nextLevelqueue
		level++
	}
	return nums
}
func levelNodes(q *list.List, level int) ([]int, *list.List) {
	queue := list.New()
	vals := make([]int, 0)
	for q.Len() != 0 {
		node := q.Front().Value.(*TreeNode)
		vals = append(vals, node.Val)
		q.Remove(q.Front())
		if level%2 == 0 {
			if node.Left != nil {
				queue.PushFront(node.Left)
			}
			if node.Right != nil {
				queue.PushFront(node.Right)
			}
		} else {
			if node.Left != nil {
				queue.PushBack(node.Left)
			}
			if node.Right != nil {
				queue.PushBack(node.Right)
			}
		}

	}
	return vals, queue
}
