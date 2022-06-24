package NodeList

/**
 * @author: xaohuihui
 * @Path: go_study/NodeList/04hasCycle.go
 * @Description:
 * @datetime: 2022/6/19 17:08:52
 * software: GoLand
**/

/* 判断环形链表 */

func hasCycle(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return false
	}
	slower, faster := head, head
	for faster != nil && faster.Next != nil {
		slower = slower.Next
		faster = faster.Next.Next
		if slower == faster {
			return true
		}
	}
	return false
}

/* 返回环形链表 环点 */
func hasCycleNode(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return nil
	}
	slower, faster := head, head
	var res *ListNode
	for faster != nil && faster.Next != nil {
		slower = slower.Next
		faster = faster.Next.Next
		if slower == faster {
			faster = head
			break
		}
	}
	for faster != slower && faster != nil {
		faster = faster.Next
		slower = slower.Next
	}
	res = faster
	return res
}