package main

import "fmt"

/**
 * @author: xaohuihui
 * @Path: go_study/algorithm/array/04max_area.go
 * @Description:
 * @datetime: 2022/6/21 16:56:16
 * software: GoLand
**/

/* 盛水最多的容器   双指针 */

func maxArea(height []int) int {
	ans := 0
	left, right := 0, len(height)-1
	for left < right {
		area := Min(height[left], height[right]) * (right - left)
		ans = Max(ans, area)
		if height[left] <= height[right] {
			left++
		} else {
			right--
		}
	}
	return ans
}

func Min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func Max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func main() {
	//height := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
	//fmt.Println(maxArea(height))
	fmt.Println(5 % 2)
	fmt.Println(6 % 2)
}
