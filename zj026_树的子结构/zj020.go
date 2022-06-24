package zj026_树的子结构

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//示例 1：
//输入：A = [1,2,3], B = [3,1]
//输出：false
//示例 2：
//输入：A = [3,4,5,1,2], B = [4,1]
//输出：true
func isSubStructure(A *TreeNode, B *TreeNode) bool {
	if A == nil || B == nil {
		return false
	}
	return recur(A, B) || isSubStructure(A.Left, B) || isSubStructure(A.Right, B)
}

func recur(A *TreeNode, B *TreeNode) bool {
	if B == nil {
		return true
	}
	if A == nil || A.Val != B.Val {
		return false
	}
	return recur(A.Left, B.Left) && recur(A.Right, B.Right)
}
