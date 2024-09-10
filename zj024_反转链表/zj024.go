package zj024_反转链表

type ListNode struct {
	Val  int
	Next *ListNode
}

// 输入: 1->2->3->4->5->NULL
// 输出: 5->4->3->2->1->NULL
// reverseList 快速的方法
func ReverseList(head *ListNode) *ListNode {
	var pre *ListNode = nil
	curr := head
	for curr != nil {
		next := curr.Next
		curr.Next = pre
		pre = curr
		curr = next
	}
	return pre
}

// reverseList1 自己写的方法
func reverseList1(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	current := head
	next := current.Next
	nextnext := current.Next.Next
	current.Next = nil
	for next != nil {
		next.Next = current
		//重新赋值
		current = next
		if nextnext != nil {
			next = nextnext
		} else {
			next = nil
		}
		if nextnext != nil {
			nextnext = nextnext.Next
		} else {
			nextnext = nil
		}
	}
	return current
}

func reverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	newHead := reverseList(head.Next)
	head.Next.Next = head
	head.Next = nil
	return newHead
}

func reverseNodeList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	newhead := reverseList(head.Next)
	head.Next.Next = newhead
	newhead.Next = nil
	return newhead
}
