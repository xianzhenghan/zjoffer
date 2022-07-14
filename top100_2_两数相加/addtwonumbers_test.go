package top100_2_两数相加

import (
	"fmt"
	"testing"
)

func TestAddTwoNumbers(t *testing.T) {
	l1 := &ListNode{
		Val: 9,
		Next: &ListNode{
			Val: 9,
			Next: &ListNode{
				Val: 9,
				Next: &ListNode{
					Val:  9,
					Next: nil,
				},
			},
		},
	}
	l2 := &ListNode{
		Val: 9,
		Next: &ListNode{
			Val:  9,
			Next: nil,
		},
	}

	relist := AddTwoNumbers(l1, l2)
	fmt.Printf("relist=%v\n", relist)
}
