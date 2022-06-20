package zj035_复杂链表的服装

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

// 方法一：hashmap，map中存的是（原节点->新节点）的映射关系
// 时间复杂度：O(n) 2次遍历
// 空间复杂度：O(n)
func copyRandomList(head *Node) *Node {
	if head == nil {
		return head
	}
	mapNode := make(map[*Node]*Node)
	// map 存储映射关系
	for cur := head; cur != nil; cur = cur.Next {
		mapNode[cur] = &Node{
			Val: cur.Val,
		}
	}
	// 	构建映射关系
	for cur := head; cur != nil; cur = cur.Next {
		mapNode[cur].Next = mapNode[cur.Next]
		mapNode[cur].Random = mapNode[cur.Random]
	}
	return mapNode[head]
}

// 原地复制之后，再拆分，空间复杂度降为O(1)
// 输入：head = [[7,null],[13,0],[11,4],[10,2],[1,0]]
// 输出：	   [[7,null],[13,0],[11,4],[10,2],[1,0]]
func copyRandomList2(head *Node) *Node {
	if head == nil {
		return head
	}
	// 1 遍历 然后复制每一个节点
	var copyNode *Node = nil
	for cur := head; cur != nil; cur = cur.Next.Next {
		copyNode = &Node{Val: cur.Val}
		copyNode.Next = cur.Next
		cur.Next = copyNode
	}

	// 2 遍历 add random 指向
	for cur := head; cur != nil; cur = cur.Next.Next {
		if cur.Random != nil {
			newNode := cur.Next
			newNode.Random = cur.Random.Next
		}

	}
	// 3 将newNode拼接
	newHead := head.Next
	for cur := head; cur != nil && cur.Next != nil; {
		temp := cur.Next     // copy节点
		cur.Next = temp.Next // 原链表
		cur = temp
	}
	return newHead
}
