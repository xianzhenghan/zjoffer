package zj006_从尾到头打印链表

import (
	"testing"
	"fmt"
)

func TestReversePrint(t *testing.T) {
	head := &ListNode{
		Val: 1111,
		Next: &ListNode{
			Val: 2222,
			Next: &ListNode{
				Val: 3333,
				Next: nil,
			},
		},
	}
	rsclice := ReversePrint(head)
	fmt.Println(rsclice)
}
