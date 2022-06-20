package zj024_反转链表

import (
	"fmt"
	"testing"
)

// TestReverseList go test -v zj024.go zj024_test.go
func TestReverseList(t *testing.T) {
	head := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val: 4,
					Next: &ListNode{
						Val:  5,
						Next: nil,
					},
				},
			},
		},
	}
	fmt.Println("List")
	printList(head)
	fmt.Println("Reverse List")
	rlist := ReverseList(head)
	printList(rlist)
}

func printList(list *ListNode) {
	current := list
	for current != nil {
		fmt.Printf("%d,", current.Val)
		current = current.Next
	}
	fmt.Println("")
}
