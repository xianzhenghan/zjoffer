package main

import (
	"fmt"
)

type BTree struct {
	value int
	left  *BTree
	right *BTree
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 递归前序遍历
func PreOrderByTraversal(root *BTree) (path []int) {
	if root == nil {
		return nil
	}
	path = make([]int, 0)
	preOrderByTraversalImp(root, &path)
	return path
}

// 递归前序遍历
func preOrderByTraversalImp(root *BTree, path *[]int) {
	if root == nil {
		return
	}
	*path = append(*path, root.value)
	if root.left != nil {
		preOrderByTraversalImp(root.left, path)
	}
	if root.right != nil {
		preOrderByTraversalImp(root.right, path)
	}
}

// 迭代方式
func preorderTraversal(root *TreeNode) []int {
	preorderSequence := make([]int, 0)
	isVisited := make(map[*TreeNode]int)
	stack := make([]*TreeNode, 0)
	stack = append(stack, root)
	isVisited[root] = 0
	for len(stack) != 0 {
		top := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if top == nil {
			continue
		}
		if isVisited[top] == 0 {
			stack = append(stack, top.Right)
			stack = append(stack, top.Left)
			stack = append(stack, top)
			isVisited[top.Right] = 0
			isVisited[top.Left] = 0
			isVisited[top] = 1
		} else {
			preorderSequence = append(preorderSequence, top.Val)
		}
	}
	return preorderSequence
}

// 迭代方式
func midorderTraversal(root *TreeNode) []int {
	preorderSequence := make([]int, 0)
	isVisited := make(map[*TreeNode]int)
	stack := make([]*TreeNode, 0)
	stack = append(stack, root)
	isVisited[root] = 0
	for len(stack) != 0 {
		top := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if top == nil {
			continue
		}
		if isVisited[top] == 0 {
			stack = append(stack, top)
			stack = append(stack, top.Right)
			stack = append(stack, top.Left)
			isVisited[top.Right] = 0
			isVisited[top.Left] = 0
			isVisited[top] = 1
		} else {
			preorderSequence = append(preorderSequence, top.Val)
		}
	}
	return preorderSequence
}

// 迭代方式
func levelOrder(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	levelOrderSequence := make([]int, 0)
	queue := make([]*TreeNode, 0)
	queue = append(queue, root)
	for len(queue) != 0 {
		size := len(queue)
		for i := 0; i < size; i++ {
			top := queue[0]
			queue = queue[1:]
			levelOrderSequence = append(levelOrderSequence, top.Val)
			if top.Left != nil {
				queue = append(queue, top.Left)
			}
			if top.Right != nil {
				queue = append(queue, top.Right)
			}
		}
	}
	return levelOrderSequence
}

func levelTraverV2(treeNode *TreeNode) []int {
	if treeNode == nil {
		return []int{}
	}
	queue := make([]*TreeNode, 0)
	levelTreeNodeTraver := make([]int, 0)
	queue = append(queue, treeNode)
	for len(queue) != 0 {
		queueSize := len(queue)
		for i := 0; i < queueSize; i++ {
			top := queue[0]
			queue = queue[1:]
			levelTreeNodeTraver = append(levelTreeNodeTraver, top.Val)
			if top.Left != nil {
				queue = append(queue, top.Left)
			}
			if top.Right != nil {
				queue = append(queue, top.Right)
			}
		}
	}
	return levelTreeNodeTraver
}

func main() {

	l := &TreeNode{6, &TreeNode{5, nil, nil}, &TreeNode{7, nil, nil}}
	r := &TreeNode{10, &TreeNode{9, nil, nil}, &TreeNode{11, nil, nil}}
	treeNode := &TreeNode{8, l, r}
	fmt.Println(levelTraverV2(treeNode))
}
