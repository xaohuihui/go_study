package NodeList

/**
 * @author: xaohuihui
 * @Path: go_study/NodeList/06getIntersectionNode.go
 * @Description:
 * @datetime: 2022/6/19 17:45:43
 * software: GoLand
**/
/* 相交链表 */

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	if headA == nil || headB == nil {
		return nil
	}
	nodeA, nodeB := headA, headB
	for nodeA != nodeB {
		if nodeA != nil {
			nodeA = nodeA.Next
		} else {
			nodeA = headB
		}

		if nodeB != nil {
			nodeB = nodeB.Next
		} else {
			nodeB = headA
		}
	}
	return nodeB
}