package zj006_从尾到头打印链表

type ListNode struct {
	Val int
	Next *ListNode
}
// ReversePrint
func ReversePrint(head *ListNode) []int {
	current := head
	sli := make([]int , 0 , 0)
	for current != nil {
		sli = append(sli,current.Val)
		current = current.Next
	}
	length := len(sli)
	rsli := make([]int , 0 ,0)
	for length > 0  {
		rsli = append(rsli,sli[length-1])
		length --
	}
	return rsli
}