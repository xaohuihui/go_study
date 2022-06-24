package NodeList

/**
 * @author: xaohuihui
 * @Path: go_study/NodeList/01palindrome_nodelist.go
 * @Description:
 * @datetime: 2022/6/19 15:15:22
 * software: GoLand
**/

/*  判断回文链表  */

func isPalindrome(head *ListNode) bool {
	left := head
	var traverse func(list *ListNode) bool
	traverse = func(right *ListNode) bool {
		if right == nil {
			return true
		}
		res := traverse(right.Next)
		res = res && (right.Val == left.Val)
		left = left.Next
		return res
	}

	return traverse(head)
}
