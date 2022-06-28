package zj018_删除链表的节点

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

func deleteNode(head *ListNode, val int) *ListNode {
	if head == nil {
		return head
	}
	if head.Val == val {
		return head.Next
	}
	next := head.Next
	curr := head
	for next != nil {
		if next.Val == val {
			curr.Next = next.Next
			next = next.Next
		} else {
			curr = next
			next = next.Next
		}
	}
	return head
}
