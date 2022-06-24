package main

/**
 * @author: xaohuihui
 * @Path: go_study/algorithm/array/03longestConsecutive.go
 * @Description:
 * @datetime: 2022/6/21 16:19:18
 * software: GoLand
**/

/* 最长连续序列   哈希表 */

func longestConsecutive(nums []int) int {
	numSet := map[int]bool{}
	// 去重数组，并标记存在的为true
	for _, num := range nums {
		numSet[num] = true
	}
	longestStreak := 0
	for num := range numSet {
		// 判断 numSet集合中是否存在当前num的前一个元素
		if !numSet[num - 1] {
			currentNum := num
			currentStreak := 1
			// 判断当前num的后一个元素是否在集合中，若存在 currentStreak加1，继续探测下一个
			for numSet[currentNum+1] {
				currentNum++
				currentStreak++
			}
			// 收集结果 找出当前元素最多的连续序列的个数
			if longestStreak < currentStreak {
				longestStreak = currentStreak
			}
		}
	}
	return longestStreak
}
