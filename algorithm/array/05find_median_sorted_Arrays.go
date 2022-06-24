package main

/**
 * @author: xaohuihui
 * @Path: go_study/algorithm/array/find_median_sorted_Arrays.go
 * @Description:
 * @datetime: 2022/6/21 17:26:05
 * software: GoLand
**/

/* 寻找两个正序数组的中位数   二分查找 */

func findMedianSortedArray(nums1, nums2 []int) float64 {
	totalLen := len(nums1) + len(nums2)
	if totalLen == 0 {
		return 0
	}
	// 总长度为奇数
	if totalLen%2 == 1 {
		midIndex := totalLen / 2
		return float64(getKthElement(nums1, nums2, midIndex+1))
	} else {
		// 总长度为偶数
		midIndex1, midIndex2 := totalLen/2-1, totalLen/2
		return float64(getKthElement(nums1, nums2, midIndex1+1)+getKthElement(nums1, nums2, midIndex2+1)) / 2.0
	}
}

func getKthElement(nums1, nums2 []int, k int) int {
	index1, index2 := 0, 0
	for {
		// 若index1超过下标范围，则中间k位置的元素存在于 nums2中
		if index1 == len(nums1) {
			return nums2[index2+k-1]
		}
		// 同理 若index2超过下标范围，则中间k位置的元素位于 nums1中
		if index2 == len(nums2) {
			return nums1[index1+k-1]
		}
		// 若k减小为 1 则 中值为 nums1 和 nums2中下一个元素的  较小小的值
		if k == 1 {
			return min(nums1[index1], nums2[index2])
		}
		// 二分查找
		half := k / 2
		newIndex1 := min(index1+half, len(nums1)) - 1
		newIndex2 := min(index2+half, len(nums2)) - 1
		pivot1, pivot2 := nums1[newIndex1], nums2[newIndex2]

		if pivot1 < pivot2 {
			k -= (newIndex1 - index1 + 1)
			index1 = newIndex1 + 1
		} else {
			k -= (newIndex2 - index2 + 1)
			index2 = newIndex2 + 1
		}
	}
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
