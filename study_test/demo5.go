package main

import "fmt"

/**
 * @author: xaohuihui
 * @Path: go_study/study_test/demo5.go
 * @Description:
 * @datetime: 2022/5/28 23:49:49
 * software: GoLand
**/
/*
go 快排
*/


func findKthLargest(nums []int, k int) int {
	nums = QuickSort(nums)
	fmt.Println(nums)
	return nums[len(nums)-k]
}

func QuickSort(nums []int) []int {
	return _quickSort(nums, 0, len(nums)-1)
}

func _quickSort(nums []int, left, right int) []int {
	if left < right {
		partitionIndex := partition(nums, left, right)
		_quickSort(nums, left, partitionIndex-1)
		_quickSort(nums, partitionIndex+1, right)
	}
	return nums
}

func partition(nums []int, left, right int) int {
	pivot := left
	index := pivot + 1

	for i := index; i <= right; i++ {
		if nums[i] < nums[pivot] {
			swap(nums, i, index)
			index += 1
		}
	}
	swap(nums, pivot, index-1)
	return index - 1
}

func swap(nums []int, i, j int) {
	nums[i], nums[j] = nums[j], nums[i]
}

func main() {
	nums := []int{2, 3, 1, 4, 6, 5}
	res := findKthLargest(nums, 2)
	fmt.Println(res)
}
