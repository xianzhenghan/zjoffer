package zj034_二叉树中和为某一值的路径

import (
	"container/list"
	"fmt"
	"slices"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//https://leetcode.cn/problems/er-cha-shu-zhong-he-wei-mou-yi-zhi-de-lu-jing-lcof/solution/go-jian-zhi-offer-34-er-cha-shu-zhong-he-de3y/

func pathSum(root *TreeNode, target int) [][]int {
	if root == nil {
		return [][]int{}
	}
	ans := [][]int{}
	var dfs func(root *TreeNode, cur int, tmp []int)
	dfs = func(root *TreeNode, cur int, tmp []int) {
		cur += root.Val
		tmp = append(tmp, root.Val)
		if root.Left == nil && root.Right == nil {
			if target == cur {
				//相当于建立临时slice再copy
				// t := make([]int, len(tmp))
				// copy(t, tmp)
				// ans = append(ans, t)
				//为什么这么做？
				//始终谨记go的参数传递为值传递，就是copy的副本
				//错误做法
				//仅仅 ans = append(ans, tmp)
				//实际上我们只需要保存每一层的slice即可
				ans = append(ans, append([]int(nil), tmp...))
				return
			}
		}
		if root.Left != nil {
			dfs(root.Left, cur, tmp)
		}
		if root.Right != nil {
			dfs(root.Right, cur, tmp)
		}
	}
	dfs(root, 0, []int{})

	return ans
}

func PathSum(root *TreeNode, target int) [][]int {
	if root == nil {
		return [][]int{}
	}
	stack := list.New()
	m := make([][]int, 0)
	dsp(root, stack, &m, target)
	return m
}

func dsp(node *TreeNode, li *list.List, m *[][]int, target int) {
	if node != nil {
		li.PushBack(node.Val)
	}
	if node.Left != nil {
		dsp(node.Left, li, m, target)
	}
	if node.Right != nil {
		dsp(node.Right, li, m, target)
	}
	// insert into []
	if node.Left == nil && node.Right == nil {
		temli := *li
		tmp := getSlice(temli, target)
		fmt.Printf("tmp=%v", tmp)
		if len(tmp) > 0 {
			*m = append(*m, tmp)
		}
	}
	li.Remove(li.Back())
}

func getSlice(li list.List, target int) []int {
	t := make([]int, 0)
	sum := 0
	for e := li.Front(); e != nil; e = e.Next() {
		t = append(t, e.Value.(int))
		sum = sum + e.Value.(int)
	}
	fmt.Printf("sum=%d,target=%d\n", sum, target)
	if sum == target {
		return t
	} else {
		return []int{}
	}
}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func pathTarget(root *TreeNode, target int) [][]int {
	if root == nil {
		return [][]int{}
	}
	var ans = make([][]int, 0)
	sli := make([]int, 0)
	return DspTarget(root, target, sli, &ans)
}

func DspTarget(root *TreeNode, target int, sli []int, ans *[][]int) [][]int {
	sli = append(sli, root.Val)
	sli1 := slices.Clone(sli)
	if root.Left == nil && root.Right == nil && getSliceSum(sli1, target) {
		*ans = append(*ans, sli1)
	}
	if root.Left != nil {
		DspTarget(root.Left, target, sli1, ans)
	}

	if root.Right != nil {
		DspTarget(root.Right, target, sli1, ans)
	}
	return *ans
}

func getSliceSum(sli []int, target int) bool {
	sum := 0
	for _, v := range sli {
		sum += v
	}
	return sum == target
}
