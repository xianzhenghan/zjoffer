package top100_2_两数相加

import "fmt"

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	var curr = &ListNode{}
	head := curr
	carry := 0
	for l1 != nil || l2 != nil {
		if l1 != nil && l2 != nil {
			node := &ListNode{
				Val:  (l1.Val + l2.Val + carry) % 10,
				Next: nil,
			}
			curr.Next = node
			curr = curr.Next
			carry = (l1.Val + l2.Val + carry) / 10
			l1 = l1.Next
			l2 = l2.Next
		}

		if l1 == nil {
			for l2 != nil {
				node := &ListNode{
					Val:  (l2.Val + carry) % 10,
					Next: nil,
				}
				curr.Next = node
				curr = curr.Next
				carry = (l2.Val + carry) / 10
				l2 = l2.Next
			}
		}

		if l2 == nil {
			for l1 != nil {
				node := &ListNode{
					Val:  (l1.Val + carry) % 10,
					Next: nil,
				}
				curr.Next = node
				curr = curr.Next
				carry = (l1.Val + carry) / 10
				l1 = l1.Next
			}
		}
	}

	if carry > 0 {
		curr.Next = &ListNode{
			Val:  carry,
			Next: nil,
		}
	}
	return head.Next
}

func AddTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	sl1, sl2 := listToslice(l1), listToslice(l2)
	i1 := 0
	i2 := 0
	carry := 0
	nums := make([]int, 0)
	fmt.Printf("sli=%v\n", sl1)
	fmt.Printf("sl2=%v\n", sl2)
	for i1 < len(sl1) || i2 < len(sl2) {
		if i1 < len(sl1) && i2 < len(sl2) {
			nums = append(nums, (sl1[i1]+sl2[i2]+carry)%10)
			carry = (sl1[i1] + sl2[i2] + carry) / 10
			i1++
			i2++
		}

		if i1 == len(sl1) {
			for i2 < len(sl2) {
				nums = append(nums, (sl2[i2]+carry)%10)
				carry = ((sl2[i2] + carry) / 10)
				i2++
			}
		}
		if i2 == len(sl2) {
			for i1 < len(sl1) {
				nums = append(nums, (sl1[i1]+carry)%10)
				carry = ((sl1[i1] + carry) / 10)
				i1++
			}
		}
	}
	if carry > 0 {
		nums = append(nums, carry)
	}
	fmt.Printf("nums=%v\n", nums)
	return sliceTolist(nums)
}

func listToslice(l1 *ListNode) []int {
	t := l1
	sli := make([]int, 0)
	for t != nil {
		sli = append(sli, t.Val)
		t = t.Next
	}
	return sli
}

func sliceTolist(nums []int) *ListNode {
	var curr = &ListNode{}
	head := curr
	for i := 0; i < len(nums); i++ {
		node := &ListNode{
			Val:  nums[i],
			Next: nil,
		}
		curr.Next = node
		curr = curr.Next
	}
	return head.Next
}
