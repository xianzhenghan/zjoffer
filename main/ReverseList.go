package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	startNode := &ListNode{0, nil}
	cur := head
	for cur != nil {
		pre := cur.Next
		// 头插入
		cur.Next = startNode.Next
		startNode.Next = cur
		//指针后移
		cur = pre
	}
	return startNode.Next
}

func reverseList2(head *ListNode) *ListNode {
	var pre *ListNode = nil
	var cur = head
	for cur != nil {
		var next = cur.Next
		cur.Next = pre
		pre = cur
		cur = next
	}
	return pre
}

func main() {
	fmt.Println(reverseList(&ListNode{1, &ListNode{2, &ListNode{3, nil}}}))
}
