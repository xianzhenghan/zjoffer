package zj025__合并两个排序的链表

type ListNode struct {
	Val  int
	Next *ListNode
}

//输入：1->2->4, 1->3->4
//输出：1->1->2->3->4->4
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	currL1 := l1
	currL2 := l2
	var head = &ListNode{}
	if currL1.Val <= currL2.Val {
		head.Next = currL1
	} else {
		head.Next = currL2
	}
	tail := head
	for currL1 != nil && currL2 != nil {
		if currL1.Val <= currL2.Val {
			tail.Next = currL1
			currL1 = currL1.Next
			tail = tail.Next
		} else {
			tail.Next = currL2
			currL2 = currL2.Next
			tail = tail.Next
		}
	}

	if currL1 != nil {
		tail.Next = currL1
	}
	if currL2 != nil {
		tail.Next = currL2
	}
	return head.Next
}
