package NodeList

/**
 * @author: xaohuihui
 * @Path: go_study/NodeList/mergeKLists.go
 * @Description:
 * @datetime: 2022/6/19 15:37:31
 * software: GoLand
**/

/* 合并k个升序链表 */

func mergeKLists(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}
	return mergeArr(lists)
}

func mergeArr(lists []*ListNode) *ListNode {
	if len(lists) <= 1 {
		return lists[0]
	}
	// 获取lists中间坐标
	index := len(lists) / 2
	left := mergeArr(lists[0:index])
	right := mergeArr(lists[index:])
	return merge(left, right)
}

func merge(l1, l2 *ListNode) *ListNode {
	if l1 == nil && l2 == nil {
		return nil
	} else if l1 != nil && l2 == nil {
		return l1
	} else if l1 == nil && l2 != nil {
		return l2
	}
	// 当l1和l2都存在时
	var newHead *ListNode
	var head *ListNode
	newHead, head = nil, nil

	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			if head == nil {
				newHead = l1
				head = l1
			} else {
				newHead.Next = l1
				newHead = newHead.Next
			}
			l1 = l1.Next
		} else {
			if head == nil {
				newHead = l2
				head = l2
			} else {
				newHead.Next = l2
				newHead = newHead.Next
			}
			l2 = l2.Next
		}
	}
	// 剩余的链表节点添加
	if l1 != nil {
		newHead.Next = l1
	} else if l2 != nil {
		newHead.Next = l2
	}
	return head
}