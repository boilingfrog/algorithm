package leetcode

import "fmt"

/*
给定两个大小为 m 和 n 的有序数组 nums1 和 nums2。

请你找出这两个有序数组的中位数，并且要求算法的时间复杂度为 O(log(m + n))。

你可以假设 nums1 和 nums2 不会同时为空。

示例 1:

nums1 = [1, 3]
nums2 = [2]

则中位数是 2.0
示例 2:

nums1 = [1, 2]
nums2 = [3, 4]

则中位数是 (2 + 3)/2 = 2.5
*/
func FindMedianSortedArrays(nums1 []int, nums2 []int) float64 {

	numAll := make([]int, 0)
	if len(nums1) > 0 {
		for _, item := range nums1 {
			numAll = append(numAll, item)
		}
	}
	if len(nums2) > 0 {
		for _, item := range nums2 {
			numAll = append(numAll, item)
		}
	}

	fmt.Println(numAll)

	return 121
}
