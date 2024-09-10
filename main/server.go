package main

import (
	"container/list"
	"fmt"
)

/**
链表中结点的数目为 sz
1 <= sz <= 30
0 <= Node.val <= 100
1 <= n <= sz
*/

type ListNode struct {
	Val  int
	Next *ListNode
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return isSubSymmetric(root.Left, root.Right)
}
func isSubSymmetric(left *TreeNode, right *TreeNode) bool {
	if left == nil || right == nil {
		return left == right
	}
	return left.Val == right.Val && isSubSymmetric(left.Left, right.Right) && isSubSymmetric(left.Right, right.Left)
}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	if root.Left == nil && root.Right == nil {
		return &TreeNode{root.Val, nil, nil}
	}
	newRoot := &TreeNode{root.Val, nil, nil}
	newLeft := invertTree(root.Left)
	newRight := invertTree(root.Right)
	newRoot.Right = newLeft
	newRoot.Left = newRight
	return newRoot
}

func invertTreeV2(root *TreeNode) *TreeNode {
	if root != nil {
		root.Left, root.Right = invertTree(root.Right), invertTree(root.Left)
	}
	return root
}

func isSubtree(root *TreeNode, subRoot *TreeNode) bool {
	if root == nil {
		return false
	}
	return isSametree(root, subRoot) || isSubtree(root.Left, subRoot) || isSubtree(root.Right, subRoot)
}

func isSametree(left *TreeNode, right *TreeNode) bool {
	if left == nil && right == nil {
		return true
	} else if (left == nil && right != nil) || (right == nil && left != nil) {
		return false
	}
	return left.Val == right.Val && isSametree(left.Left, right.Left) && isSametree(left.Right, right.Right)
}

/*
*
Input: list1 = [1,2,4], list2 = [1,3,4]
Output: [1,1,2,3,4,4]
Example 2:

Input: list1 = [], list2 = []
Output: []
Example 3:

Input: list1 = [], list2 = [0]
Output: [0]
*/
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	if list1 == nil {
		return list2
	}
	if list2 == nil {
		return list1
	}
	head := &ListNode{0, nil}
	cur := head
	var next *ListNode = nil
	for list1 != nil && list2 != nil {
		if list1.Val <= list2.Val {
			next = &ListNode{list1.Val, nil}
			list1 = list1.Next
		} else {
			next = &ListNode{list2.Val, nil}
			list2 = list2.Next
		}
		cur.Next = next
		cur = next
	}
	if list1 == nil && list2 == nil {
		return head.Next
	}

	if list1 != nil {
		cur.Next = list1
	} else {
		cur.Next = list2
	}
	return head.Next
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	size := 0
	var pre = &ListNode{}
	pre = head
	for pre != nil {
		size++
		pre = pre.Next
	}
	fmt.Println(size)
	dummy := &ListNode{0, head}
	cur := dummy
	for i := 0; i < (size - n); i++ {
		cur = cur.Next
	}
	cur.Next = cur.Next.Next
	return dummy.Next
}

/*
*
递归
*/
func revserListV1(head *ListNode) *ListNode {
	return nil
}

/*
*
迭代
*/
func revserListV2(head *ListNode) *ListNode {
	pre := &ListNode{}
	current := head
	next := &ListNode{}
	for current != nil {
		next = current.Next
		current.Next = pre
		pre = current
		current = next
	}
	return pre
}

func helper(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	nextResult := helper(head.Next)
	head.Next.Next = head
	head.Next = nil
	return nextResult
}
func printLinkList(head *ListNode) {
	dummy := &ListNode{
		Val:  0,
		Next: head,
	}
	cur := dummy
	for cur.Next != nil && cur != nil {
		fmt.Print(cur.Next)
		cur = cur.Next
	}
}

/*
*
m == matrix.length
n == matrix[i].length
1 <= m, n <= 10
-100 <= matrix[i][j] <= 100
*/
func spiralOrder(matrix [][]int) []int {
	fmt.Println(matrix)
	m := len(matrix)
	n := len(matrix[0])
	visited := make([][]int, m)
	fmt.Println("m=", m, " n=", n)
	for i := 0; i < m; i++ {
		visited[i] = make([]int, n)
	}
	result := make([]int, 0)
	dspPrint(matrix, visited, &result, 0, 0, m, n, false)
	fmt.Println("result=", result)
	return result
}

func dspPrint(matrix [][]int, visited [][]int, result *[]int, i, j int, m, n int, isUp bool) {
	fmt.Println(visited)

	if i >= 0 && i < m && j >= 0 && j < n && visited[i][j] == 0 {
		*result = append(*result, matrix[i][j])
		visited[i][j] = 1
	}
	if isUp {
		if i-1 >= 0 && i-1 < m && j >= 0 && j < n && visited[i-1][j] == 0 {
			dspPrint(matrix, visited, result, i-1, j, m, n, true)
		}
	}
	// 往右
	if i >= 0 && i < m && j+1 >= 0 && j+1 < n && visited[i][j+1] == 0 {
		dspPrint(matrix, visited, result, i, j+1, m, n, false)
	}

	// 往下
	if i+1 >= 0 && i+1 < m && j >= 0 && j < n && visited[i+1][j] == 0 {
		dspPrint(matrix, visited, result, i+1, j, m, n, false)
	}

	// 往左
	if i >= 0 && i < m && j-1 >= 0 && j-1 < n && visited[i][j-1] == 0 {
		dspPrint(matrix, visited, result, i, j-1, m, n, false)
	}
	// 往上
	if i-1 >= 0 && i-1 < m && j >= 0 && j < n && visited[i-1][j] == 0 {
		dspPrint(matrix, visited, result, i-1, j, m, n, true)
	}
}

/*
*
Example 1:

Input: pushed = [1,2,3,4,5], popped = [4,5,3,2,1]
Output: true
Explanation: We might do the following sequence:
push(1), push(2), push(3), push(4),
pop() -> 4,
push(5),
pop() -> 5, pop() -> 3, pop() -> 2, pop() -> 1
Example 2:

Input: pushed = [1,2,3,4,5], popped = [4,3,5,1,2]
Output: false
Explanation: 1 cannot be popped before 2.

Constraints:

1 <= pushed.length <= 1000
0 <= pushed[i] <= 1000
All the elements of pushed are unique.
popped.length == pushed.length
popped is a permutation of pushed.
*/
func validateStackSequences(pushed []int, popped []int) bool {
	stack := list.List{}
	popIdx := 0
	for i := 0; i < len(pushed); {
		if pushed[i] == popped[popIdx] {
			popIdx++
			i++
			continue
		}
		if stack.Len() != 0 {
			elem := stack.Front()
			if elem.Value.(int) == popped[popIdx] {
				stack.Remove(elem)
				popIdx++
			}
		} else {
			stack.PushFront(pushed[i])
			i++
		}
	}
	return popIdx == len(popped) && stack.Len() == 0
}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}
	stackSlice := make([]*TreeNode, 0)
	stackSlice = append(stackSlice, root)
	res := make([][]int, 0)
	for len(stackSlice) != 0 {
		tem := make([]int, 0)
		length := len(stackSlice)
		for i := 0; i < length; i++ {
			tem = append(tem, stackSlice[0].Val)
			if stackSlice[0].Left != nil {
				stackSlice = append(stackSlice, stackSlice[0].Left)
			}
			if stackSlice[0].Right != nil {
				stackSlice = append(stackSlice, stackSlice[0].Right)
			}
			stackSlice = stackSlice[1:]
		}
		res = append(res, tem)
	}
	return res
}

func permute(nums []int) [][]int {
	dp := make([][][]int, len(nums))
	dp[0] = [][]int{{nums[0]}}
	for i := 1; i < len(nums); i++ {
		dpi := dp[i]
		temdpi := make([][]int, 0)
		for _, co := range dpi {
			for pos := 0; pos <= len(co); pos++ {
				sli := make([]int, 0)
				sli = append(sli, co[0:pos]...)
				sli = append(sli, nums[i])
				sli = append(sli, co[pos:]...)
				temdpi = append(temdpi, sli)
			}
		}
		dp[i] = temdpi
	}
	return dp[len(nums)-1]
}

func permute2(nums []int) [][]int {
	var result [][]int

	var backtrack func([]int, []int)
	backtrack = func(nums []int, path []int) {
		if len(nums) == 0 {
			result = append(result, append([]int(nil), path...))
			return
		}
		for i := 0; i < len(nums); i++ {
			newNums := append([]int(nil), nums[:i]...)
			newNums = append(newNums, nums[i+1:]...)
			newPath := append([]int(nil), path...)
			newPath = append(newPath, nums[i])
			backtrack(newNums, newPath)
		}
	}

	backtrack(nums, []int{})
	return result
}

func main() {
	fmt.Println(permute([]int{1, 2}))
}
