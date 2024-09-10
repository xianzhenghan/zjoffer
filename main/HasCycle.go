package main

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

func hasCycle(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return false
	}
	slow, fast := head, head.Next
	for fast != slow {
		if fast == nil || fast.Next == nil {
			return false
		}
		slow = slow.Next
		fast = fast.Next.Next
	}
	return true
}

func hasCycle2(head *ListNode) bool {
	if head == nil || head.Next == nil || head.Next.Next == nil {
		return false
	}
	slow, fast := ListNode{0, head}.Next, ListNode{0, head}.Next.Next
	for slow != nil && fast != nil {
		if slow == fast {
			return true
		}
		if slow.Next == nil || fast.Next == nil || fast.Next.Next == nil {
			return false
		}
		slow = slow.Next
		fast = fast.Next.Next
	}
	return false
}

func main() {
	dataChild := &ListNode{3, nil}
	data := &ListNode{2, dataChild}
	head := &ListNode{1, data}

	dataChild.Next = head
	fmt.Println(hasCycle2(head))
}
