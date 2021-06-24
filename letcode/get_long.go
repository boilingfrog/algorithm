package letcode

import (
	"fmt"
	"github.com/mohae/deepcopy"
	"strings"
)

func digui(mappBase map[int]int, paiArr []int) []int {

	num := len(paiArr)
	if _, ok := mappBase[paiArr[num-1]+paiArr[num-2]]; ok {
		paiArr = append(paiArr, paiArr[num-1]+paiArr[num-2])
		delete(mappBase, paiArr[num-1]+paiArr[num-2])
		paiArr = digui(mappBase, paiArr)
	}
	return paiArr
}

func getLong(nums []int) string {
	// 对nums先排序，并且放到一个hashmap中，此处省略排序的
	// 此处放入省略
	mappBase := make(map[int]int)
	mappBase[10] = 0
	mappBase[20] = 1
	mappBase[30] = 2
	mappBase[40] = 3
	mappBase[50] = 4

	var paiAll = make([]string, 0)
	num := len(nums)
	//两层循环
	for i := 0; i < num; i++ {

		for j := i + 1; j < num; j++ {
			// deep copy原始的hashmap
			mappBaseCheck := deepcopy.Copy(mappBase)
			mappBaseCheckRes := mappBaseCheck.(map[int]int)
			delete(mappBaseCheckRes, nums[i])
			//每次放入递归数组中两个初始值
			var paiArr = make([]int, 0)
			paiArr = append(paiArr, nums[i])
			paiArr = append(paiArr, nums[j])
			// 删除hashmap中的已经取出的值
			delete(mappBaseCheckRes, nums[j])
			//递归
			ress := digui(mappBaseCheckRes, paiArr)
			// 结果转成string,放到数组中
			proString := strings.Replace(strings.Trim(fmt.Sprint(ress), "[]"), " ", ",", -1)
			paiAll = append(paiAll, proString)
		}

	}
	check := 0
	lenn := 0
	// 循环数组找到最长的字符串
	for key, item := range paiAll {
		if len(item) > lenn {
			check = key
			lenn = len(item)
		}
	}
	return paiAll[check]
}
