package main

import "fmt"

/**
 * @author: xaohuihui
 * @Path: go_study/algorithm/array/02find_length_of_LCIS.go
 * @Description:
 * @datetime: 2022/6/21 15:38:58
 * software: GoLand
**/

/* 最长连续递增序列  快慢指针 */

func findLengthOfLCIS(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	left, right := 0, 1
	gMaxLen, maxLen := 1, 1
	for right < len(nums) {
		if nums[right] > nums[left] {
			maxLen++
		} else {
			maxLen = 1
		}
		left++
		right++
		gMaxLen = max(gMaxLen, maxLen)
	}
	return gMaxLen
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func main() {
	nums := []int{2,2,2,2,2}
	fmt.Println(findLengthOfLCIS(nums))
}
