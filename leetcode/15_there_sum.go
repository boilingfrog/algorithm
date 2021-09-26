package leetcode

import (
	"fmt"
	"strconv"
)

func ThreeSum(nums []int) [][]int {
	MergeSort(nums, len(nums))
	var res = [][]int{}
	mapp := make(map[string]int)
	for t := 0; t < (len(nums) - 2); t++ {
		target := -nums[t]
		i := t + 1
		j := len(nums) - 1
		for {
			if i >= j {
				break
			}
			// 双指针进行判断
			if (nums[i] + nums[j]) > target {
				//nowData = nums[j-1] * (j - i - 1)
				j--
			} else if (nums[i] + nums[j]) < target {
				//nowData = nums[i] * (j - i - 1)
				i++
			} else if (nums[i] + nums[j]) == target {
				//fmt.Println("henghengheng", nums[i], nums[j], nums[0])
				var resNow = []int{nums[i], nums[j], nums[t]}
				var str = strconv.Itoa(nums[i]) + strconv.Itoa(nums[j]) + strconv.Itoa(nums[t])
				if _, ok := mapp[str]; ok {
					//存在
					fmt.Println("heng")
				} else {
					res = append(res, resNow)
					mapp[str] = 1
				}
				j--
				i++
			}
		}
	}

	return res
}

// 归并排序
func MergeSort(arr []int, n int) {

	mergeSort(arr, 0, n-1)
}

// 递归使用归并排序，对arr[1.......r]的范围进行排序
func mergeSort(arr []int, l int, r int) {
	if l >= r {
		return
	}
	mid := (l + r) / 2 //问题点可能这个地方会出现内存的溢出
	mergeSort(arr, l, mid)
	mergeSort(arr, mid+1, r)
	if arr[mid] > arr[mid+1] {
		Merge(arr, l, mid, r)
	}
}

// 将arr[l......mid]和arr[mid+1.....r]两部分进行归并
func Merge(arr []int, l int, mid int, r int) {
	num := r - l + 1
	aux := make([]int, num)
	for i := l; i <= r; i++ {
		aux[i-l] = arr[i] //l到r之间，放到临时的数组空间，0到l之间有一个l的偏移量
	}
	i := l
	j := mid + 1
	for k := l; k <= r; k++ {
		if i > mid {
			arr[k] = aux[j-l]
			j++
		} else if j > r {
			arr[k] = aux[i-l]
			i++
		} else if aux[i-l] < aux[j-l] {
			arr[k] = aux[i-l]
			i++
		} else {
			arr[k] = aux[j-l]
			j++
		}
	}
}
