package leetcode

/*
假设有一个很长的花坛，一部分地块种植了花，另一部分却没有。可是，花不能种植在相邻的地块上，它们会争夺水源，两者都会死去。

给你一个整数数组flowerbed 表示花坛，由若干 0 和 1 组成，其中 0 表示没种植花，1 表示种植了花。另有一个数n ，能否在不打破种植规则的情况下种入n朵花？能则返回 true ，不能则返回 false。



示例 1：

输入：flowerbed = [1,0,0,0,1], n = 1
输出：true


示例 2：

输入：flowerbed = [1,0,0,0,1], n = 2
输出：false

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/can-place-flowers
*/

/*
思路

花坛中的连续区域分为三种，最左边的连续，中间的连续，和最右边的连续

维护 begin 表示上一朵已经种植的花的下标位置，初始时 begin=−1，表示尚未遇到任何已经种植的花。

从左往右遍历数组 flowerbed，当遇到 flowerbed[i]=1 时根据 begin 和 i 的值计算上一个区间内可以种植花的最多数量，然后令 begin=i，继续遍历数组 flowerbed 剩下的元素。

遍历数组 flowerbed 结束后，根据数组 begin 和长度 mm 的值计算最后一个区间内可以种植花的最多数量。

判断整个花坛内可以种入的花的最多数量是否大于或等于 nn。

*/
func canPlaceFlowers(flowerbed []int, n int) bool {
	le := len(flowerbed)
	if le == 0 {
		return false
	}

	nums := 0
	begin := -1
	for i := 0; i < le; i++ {
		if flowerbed[i] == 1 {
			if begin == -1 {
				nums += i / 2
			} else {
				nums += (i - begin - 2) / 2
			}
			begin = i
		}
	}

	if begin == -1 {
		nums += (le + 1) / 2
	} else {
		nums += (le - begin - 1) / 2
	}

	return nums >= n
}
