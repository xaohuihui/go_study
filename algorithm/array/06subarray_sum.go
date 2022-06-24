package main

/**
 * @author: xaohuihui
 * @Path: go_study/algorithm/array/06subarray_sum.go
 * @Description:
 * @datetime: 2022/6/21 19:01:12
 * software: GoLand
**/
/* 和为k的子数组个数   前缀和 + 哈希表 */

func subarraySum(nums []int, k int) int {
	count, pre := 0, 0
	m := map[int]int{}
	m[0] = 1
	for i := 0; i < len(nums); i++ {
		pre += nums[i]
		// 若前缀和 - k 存在哈希表中， 则该组合成立 count 加上该前缀和 - k 的元素个数 为该前缀树
		if _, ok := m[pre - k]; ok {
			count += m[pre - k]
		}
		m[pre] += 1
	}
	return count
}
