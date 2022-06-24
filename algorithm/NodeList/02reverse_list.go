package NodeList

import "fmt"

/**
 * @author: xaohuihui
 * @Path: go_study/NodeList/reverse_list.go
 * @Description:
 * @datetime: 2022/6/19 15:31:38
 * software: GoLand
**/

/* 反转链表 */

//type ListNode struct {
//	Val  int
//	Next *ListNode
//}


func reverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	// 递归将head下面节点压栈
	last := reverseList(head.Next)
	// head 下个节点 指针指向 head
	head.Next.Next = head
	// head 本节点指针指向nil
	head.Next = nil
	return last
}

func reverseList2(head *ListNode) *ListNode {
	var prev *ListNode
	curr := head
	nxt := head
	for curr != nil {
		nxt = curr.Next
		curr.Next = prev

		prev = curr
		curr = nxt
	}
	return prev
}

func main() {
	node5 := &ListNode{Val: 1, Next: nil}
	node4 := &ListNode{Val: 2, Next: node5}
	node3 := &ListNode{Val: 3, Next: node4}
	node2 := &ListNode{Val: 4, Next: node3}
	node1 := &ListNode{Val: 5, Next: node2}
	head := reverseList2(node1)
	for head != nil {
		fmt.Printf("%d", head.Val)
		head = head.Next
	}
}