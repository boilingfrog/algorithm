package letcode

/*
实现获取 下一个排列 的函数，算法需要将给定数字序列重新排列成字典序中下一个更大的排列。

如果不存在下一个更大的排列，则将数字重新排列成最小的排列（即升序排列）。

必须 原地 修改，只允许使用额外常数空间。



示例 1：

输入：nums = [1,2,3]
输出：[1,3,2]

示例 2：

输入：nums = [3,2,1]
输出：[1,2,3]

示例 3：

输入：nums = [1,1,5]
输出：[1,5,1]

示例 4：

输入：nums = [1]
输出：[1]


提示：

1 <= nums.length <= 100
0 <= nums[i] <= 100


思路：其实这道题如果不知道规律的话，单纯看样例真的是一脸懵逼。根据disguss中的解释，next permutation的意思如下：

首先从数组右向左查找，找到第一个满足nums[k]<nums[k+1]的下标k，如果不存在这个k，那么这个数组是一个递减数组，直接返回数组的逆序。

然后找到下标l满足：l>k并且nums[l]>nums[k]，交换nums[k]和nums[l]

最后对大于k的子数组进行逆序排序

*/

func NextPermutation(nums []int) {
	i, j := 0, 0
	for i = len(nums) - 2; i >= 0; i-- {
		if nums[i] < nums[i+1] {
			break
		}
	}
	if i >= 0 {
		for j = len(nums) - 1; j > i; j-- {
			if nums[j] > nums[i] {
				break
			}
		}
		swap(&nums, i, j)
	}
	reverse(&nums, i+1, len(nums)-1)
}

func reverse(nums *[]int, i, j int) {
	for i < j {
		swap(nums, i, j)
		i++
		j--
	}
}

func swap(nums *[]int, i, j int) {
	(*nums)[i], (*nums)[j] = (*nums)[j], (*nums)[i]
}
