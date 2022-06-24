package main

/**
 * @author: xaohuihui
 * @Path: go_study/study_test/demo6.go
 * @Description:
 * @datetime: 2022/5/30 9:52:04
 * software: GoLand
**/

/* 获取二叉树的宽度 */

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type element struct {
	root  *TreeNode
	depth int
	pos   int
}

func widthOfBinaryTree(root *TreeNode) int {
	queue := make([]*element, 0, 0)
	//queue := []*element{&element{root: root, depth: 0, pos: 0}}
	queue = append(queue, &element{root: root, depth: 0, pos: 0})
	cur_depth, left, ans, n := 0, 0, 0, len(queue)
	for i := 0; i < n; i++ {
		ele := queue[i]
		if ele.root != nil {
			queue = append(queue, &element{ele.root.Left, ele.depth + 1, ele.pos * 2})
			queue = append(queue, &element{ele.root.Right, ele.depth + 1, ele.pos*2 + 1})
			if cur_depth != ele.depth {
				cur_depth = ele.depth
				left = ele.pos
			}
			ans = max(ele.pos-left+1, ans)
			n = len(queue)
		}
	}
	return ans
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func main() {

}
