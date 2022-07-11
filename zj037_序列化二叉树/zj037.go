package zj037_序列化二叉树

import (
	"strconv"
	"strings"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/**
 * 代码中的类名、方法名、参数名已经指定，请勿修改，直接返回方法规定的值即可
 *
 * @param root TreeNode类
 * @return TreeNode类
 */
func dfs(root *TreeNode, str *string) {
	if root == nil {
		*str += "null,"
		return
	}
	*str += strconv.Itoa(root.Val) + ","
	dfs(root.Left, str)
	dfs(root.Right, str)
	return
}
func Serialize(root *TreeNode) string {
	// write code her
	if root == nil {
		return ""
	}
	ans := ""
	dfs(root, &ans)
	return ans
}
func Deserialize(s string) *TreeNode {
	// write code here
	if len(s) == 0 {
		return nil
	}
	strs := strings.Split(s, ",")
	return de_dfs(&strs)
}
func de_dfs(str *[]string) *TreeNode {
	if (*str)[0] == "null" {
		*str = (*str)[1:] //没办法go是值传递，你需要改变切片的长度，让构造可以进行下去
		return nil
	}
	nums, _ := strconv.Atoi((*str)[0])
	*str = (*str)[1:]
	root := &TreeNode{Val: nums}
	if len(*str) != 0 {
		root.Left = de_dfs(str)
	}
	if len(*str) != 0 {
		root.Right = de_dfs(str)
	}
	return root
}
