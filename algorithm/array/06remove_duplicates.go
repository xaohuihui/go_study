package main

/**
 * @author: xaohuihui
 * @Path: go_study/algorithm/array/06remove_duplicates.go
 * @Description:
 * @datetime: 2022/6/21 18:27:25
 * software: GoLand
**/

/* 删除有序数组中的重复项 快慢指针 */

func removeDuplicates(nums []int) int {
	if len(nums) <= 1 {
		return len(nums)
	}
	n := len(nums)
	slower := 1
	for faster := 1; faster < n; faster++ {
		// 追齐快慢指针
		if nums[faster] != nums[faster-1] {
			nums[slower] = nums[faster]
			slower++
		}
	}
	return slower
}
