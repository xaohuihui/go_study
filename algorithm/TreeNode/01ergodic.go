package main

import "fmt"

/**
 * @author: xaohuihui
 * @Path: go_study/algorithm/TreeNode/01ergodic.go
 * @Description:
 * @datetime: 2022/6/23 10:46:55
 * software: GoLand
**/

// 遍历方式

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// PreOrderRecursion 二叉树递归方式先序遍历
func PreOrderRecursion(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	res := append([]int{root.Val}, PreOrderRecursion(root.Left)...)
	res = append(res, PreOrderRecursion(root.Right)...)
	return res
}

// PreOrderTraverse 前序遍历  迭代
func PreOrderTraverse(root *TreeNode) []int {
	// 栈用来存放节点
	stack := []*TreeNode{}
	node := root
	res := []int{}
	for node != nil || len(stack) > 0 {
		for node != nil {
			// 将根和左子树加入到结果集合中，与 stack栈中
			res = append(res, node.Val)
			stack = append(stack, node)
			node = node.Left
		}
		// 出栈并取右子树
		node = stack[len(stack)-1].Right
		stack = stack[:len(stack)-1]
	}
	return res
}

// MidOrderTraverse1 递归方式实现中序遍历
func MidOrderTraverse1(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	res := append(MidOrderTraverse1(root.Left), root.Val)
	res = append(res, MidOrderTraverse1(root.Right)...)
	return res
}

// MidOrderTraverse1 递归方式实现中序遍历
func MidOrderTraverse2(root *TreeNode) []int {
	// 非递归实现
	stack := []*TreeNode{}
	node := root
	res := []int{}

	for node != nil || len(stack) > 0 {
		for node != nil {
			stack = append(stack, node)
			node = node.Left
		}

		node = stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		res = append(res, node.Val)
		node = node.Right
	}
	return res
}

// 二叉树后续遍历  递归
func postorderTraverse1(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	res := append(postorderTraverse1(root.Left), postorderTraverse1(root.Right)...)
	res = append(res, root.Val)
	return res
}

// 二叉树后续遍历  迭代
func postorderTraverse2(root *TreeNode) []int {
	stack := []*TreeNode{}
	var prev *TreeNode
	res := []int{}
	for root != nil || len(stack) > 0 {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}
		root = stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		if root.Right == nil || root.Right == prev {
			res = append(res, root.Val)
			prev = root
			root = nil
		} else {
			stack = append(stack, root)
			root = root.Right
		}
	}
	return res
}

func main() {
	root4 := &TreeNode{Val: 4, Left: nil, Right: nil}
	root5 := &TreeNode{Val: 5, Left: nil, Right: nil}
	root6 := &TreeNode{Val: 6, Left: nil, Right: nil}
	root2 := &TreeNode{Val: 2, Left: root4, Right: root5}
	root3 := &TreeNode{Val: 3, Left: nil, Right: root6}
	root1 := &TreeNode{Val: 1, Left: root2, Right: root3}
	fmt.Println(postorderTraverse2(root1))
}
