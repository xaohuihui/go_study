package main

/**
 * @author: xaohuihui
 * @Path: go_study/algorithm/string/01max_palindrome_str.go
 * @Description:
 * @datetime: 2022/6/20 10:00:07
 * software: GoLand
**/

/* 最长回文子串 */

func longestPalindrome(s string) string {
	if len(s) <= 1 {
		return s
	}
	maxRes, maxStr := 0, ""
	for i := 0; i < len(s); i++ {
		str1 := palindrome(s, i, i)
		str2 := palindrome(s, i, i+1)
		if len(str1) > maxRes {
			maxStr = str1
			maxRes = len(str1)
		}
		if len(str2) > maxRes {
			maxStr = str2
			maxRes = len(str2)
		}
	}
	return maxStr
}

func palindrome(s string, l, r int) string {
	for l >= 0 && r < len(s) && s[l] == s[r] {
		l--
		r++
	}
	return s[l+1 : r]
}
