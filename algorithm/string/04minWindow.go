package main

/**
 * @author: xaohuihui
 * @Path: go_study/algorithm/string/04minWindow.go
 * @Description:
 * @datetime: 2022/6/21 10:28:25
 * software: GoLand
**/
/* 最小覆盖子串 */

func minWindow(s, t string) string {
	need := make(map[byte]int)
	window := make(map[byte]int)

	for i := range t {
		need[t[i]]++
	}

	left, right := 0, 0
	valid := 0
	// t 去重后的数量
	t_set_n := len(need)
	// 初始化为最大长度 + 1
	minLen := len(s) + 1
	minStr := ""
	for right < len(s) {
		d := s[right]
		right++
		window[d]++
		// 判断 t中的该字符和 windows的该字符的数量是否相同，若遇到重复的字符，前面的统计不会使valid增加
		if need[d] == window[d] {
			valid++
		}
		// 判断valid是否和t的去重后长度相同, 然后左指针向后移动，找到最短字符串不满足前面条件跳出
		for valid == t_set_n {
			if right - left < minLen {
				minLen = right - left
				minStr = s[left:right]
			}
			c := s[left]
			left++
			window[c]--
			if  window[c] < need[c] {
				valid--
			}
		}
	}
	return minStr
}
