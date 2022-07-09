package zj031_栈的压入弹出序列

import (
	"container/list"
	"fmt"
)

func validateStackSequences(pushed []int, popped []int) bool {
	if len(pushed) == 0 {
		return true
	}
	stack := list.New()
	i := 0
	for _, v := range pushed {
		stack.PushBack(v)
		for stack.Len() > 0 && stack.Back().Value.(int) == popped[i] {
			tmp := stack.Back()
			stack.Remove(tmp)
			i++
		}
	}
	return stack.Len() == 0
}

func validateStackSequences2(pushed []int, popped []int) bool {
	if len(pushed) == 0 {
		return true
	}
	stack := list.New()
	i := 0
	j := 0
	for (j < len(popped) && i < len(pushed)) || stack.Len() > 0 {
		fmt.Printf("i=%d,j=%d", i, j)
		if i < len(pushed) {
			if pushed[i] == popped[j] {
				i++
				j++
			} else {
				if stack.Len() > 0 {
					num := stack.Back()
					if num.Value.(int) == popped[j] {
						stack.Remove(num)
						j++
					} else {
						stack.PushBack(pushed[i])
						i++
					}
				} else {
					stack.PushBack(pushed[i])
					i++
				}
			}
		} else {
			num := stack.Back()
			if num.Value.(int) == popped[j] {
				stack.Remove(num)
				j++
			} else {
				return false
			}
		}
	}
	return true
}
