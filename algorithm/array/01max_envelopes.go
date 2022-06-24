package main

import (
	"fmt"
	"sort"
)

/**
 * @author: xaohuihui
 * @Path: go_study/algorithm/array/01max_envelopes.go
 * @Description:
 * @datetime: 2022/6/21 14:16:32
 * software: GoLand
**/

/* 俄罗斯套娃信封问题  基于二分查找的 动态规划 */

func maxEnvelopes(envelopes [][]int) int {
	// 先将第一维度 宽度排序  底层使用快排算法
	sort.Slice(envelopes, func(i, j int) bool {
		a, b := envelopes[i], envelopes[j]
		// 题目要求 宽度和高度都需要比前面的大的时候才能套在一起 所以 后面条件才是大于
		return a[0] < b[0] || a[0] == b[0] && a[1] > b[1]
	})
	fmt.Println(envelopes)
	// f 保存 符合递进的高度
	f := make([]int, 0, 0)
	for _, e := range envelopes {
		h := e[1]
		// 二分查找
		if i := sort.SearchInts(f, h); i < len(f) {
			f[i] = h
		} else {
			f = append(f, h)
		}
	}
	return len(f)
}

func main() {
	envelopes := [][]int{{4, 5}, {4, 6}, {6, 7}, {2, 3}, {1, 1}}
	fmt.Println(maxEnvelopes(envelopes))
}
