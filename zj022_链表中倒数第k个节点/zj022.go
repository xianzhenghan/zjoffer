package zj022_链表中倒数第k个节点

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

//给定一个链表: 1->2->3->4->5, 和 k = 2.
//返回链表 4->5.
func getKthFromEnd(head *ListNode, k int) *ListNode {
	if head == nil {
		return nil
	}
	length := 0
	current := head
	for current != nil {
		length++
		current = current.Next
	}
	fmt.Printf("length=%d", length)
	prenum := length - k
	current = head
	for prenum > 0 {
		current = current.Next
		prenum--
	}
	return current.Next
}

// 快慢指针
func getKthFromEnd2(head *ListNode, k int) *ListNode {
	fast, slow := head, head
	for fast != nil && k > 0 {
		fast = fast.Next
		k--
	}
	for fast != nil {
		fast = fast.Next
		slow = slow.Next
	}
	return slow
}
