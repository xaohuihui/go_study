package main

import "fmt"

/**
 * @author: xaohuihui
 * @Path: go_study/algorithm/string/03.go
 * @Description:
 * @datetime: 2022/6/21 9:47:14
 * software: GoLand
**/

/* 无重复字符的最长子串 */

func lengthOfLongestSubstring(s string) int {
	tmp_map := make(map[byte]int)
	right, res, n := -1, 0, len(s)
	for i := 0; i < n; i++ {
		// 左指针向右移动一格
		if i != 0 {
			delete(tmp_map, s[i-1])
		}
		// 右指针持续向右移动
		fmt.Println(tmp_map)
		for right+1 < n && tmp_map[s[right+1]] == 0 {
			tmp_map[s[right+1]]++
			right++
			fmt.Println("ssssss", tmp_map)
		}
		res = max(res, right+1)
	}
	return res
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func main() {
	s := "asdfghjklazxs"
	fmt.Println(lengthOfLongestSubstring(s))
}