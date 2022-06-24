package main

import (
	"math"
)

/**
 * @author: xaohuihui
 * @Path: go_study/algorithm/string/02longest_common_prefix.go
 * @Description:
 * @datetime: 2022/6/20 10:51:22
 * software: GoLand
**/

/* 最长公共前缀 */

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	first := strs[0]
	if first == "" {
		return ""
	}
	minLen := math.MaxInt64
	for i := 1; i < len(strs); i++ {
		ilen := twoStrLongestCommonPrefix(first, strs[i])
		minLen = min(ilen, minLen)
	}
	return first[0:minLen]
}

func twoStrLongestCommonPrefix(s, tmp string) int {
	i, j := 0, 0
	cnt := 0
	for i < len(s) && j < len(tmp) {
		if s[i] == tmp[j] {
			cnt++
		} else {
			return cnt
		}
		i++
		j++
	}
	return cnt
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
