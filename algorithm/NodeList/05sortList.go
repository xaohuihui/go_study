package NodeList

import "fmt"

/**
 * @author: xaohuihui
 * @Path: go_study/NodeList/05sortList.go
 * @Description:
 * @datetime: 2022/6/19 17:31:38
 * software: GoLand
**/
/* 排序链表 */

func sortList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	return mergeSort(head)
}

func mergeSort(head *ListNode) *ListNode {
	if head.Next != nil {
		slower := getCenter(head)
		nxt := slower.Next
		slower.Next = nil
		left := mergeSort(head)
		right := mergeSort(nxt)
		head = mergeList(left, right)
	}
	return head
}

func mergeList(left, right *ListNode) *ListNode {
	var newHead *ListNode
	var head *ListNode
	newHead, head = nil, nil
	for left != nil && right != nil {
		if left.Val < right.Val {
			if head == nil {
				newHead = left
				head = left
			} else {
				newHead.Next = left
				newHead = newHead.Next
			}
			left = left.Next
		} else {
			if head == nil {
				newHead = right
				head = right
			} else {
				newHead.Next = right
				newHead = newHead.Next
			}
			right = right.Next
		}
	}
	if left != nil {
		newHead.Next = left
	}
	if right != nil {
		newHead.Next = right
	}
	return head
}

func getCenter(head *ListNode) *ListNode {
	slower, faster := head, head.Next
	for faster != nil && faster.Next != nil {
		slower = slower.Next
		faster = faster.Next.Next
	}
	return slower
}

func main() {
	node5 := &ListNode{Val: 1, Next: nil}
	node4 := &ListNode{Val: 2, Next: node5}
	node3 := &ListNode{Val: 3, Next: node4}
	node2 := &ListNode{Val: 4, Next: node3}
	node1 := &ListNode{Val: 5, Next: node2}
	l := sortList(node1)
	for l != nil {
		fmt.Println(l.Val)
		l = l.Next
	}
}