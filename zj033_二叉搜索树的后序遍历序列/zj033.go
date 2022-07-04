package zj033_二叉搜索树的后序遍历序列

func verifyPostorder(postorder []int) bool {
	if len(postorder) < 1 {
		return true
	}
	root_value := postorder[len(postorder)-1]
	index := len(postorder) - 1
	for idx := 0; idx < len(postorder)-1; idx++ {
		if index == len(postorder)-1 && postorder[idx] > root_value {
			index = idx
		}
		if index != len(postorder)-1 && postorder[idx] < root_value {
			return false
		}
	}
	return verifyPostorder(postorder[0:index]) && verifyPostorder(postorder[index:len(postorder)-1])
}

//
func getBGvalue(value int, postorder []int) int {
	for idx := 0; idx < len(postorder); idx++ {
		if postorder[idx] > value {
			return idx
		}
	}
	return len(postorder)
}

func isRightBGLeft(left []int, right []int, rv int) bool {
	if len(right) == 0 || len(left) == 0 {
		return true
	}

	for l := 0; l < len(left); l++ {
		if left[l] > rv {
			return false
		}
	}
	for r := 0; r < len(right); r++ {
		if right[r] < rv {
			return false
		}
	}

	for r := 0; r < len(right); r++ {
		for l := 0; l < len(left); l++ {
			if left[l] > right[r] {
				return false
			}
		}
	}
	return true
}
