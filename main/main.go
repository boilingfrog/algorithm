package main

import (
	"algorithm/letcode"
	"fmt"
	"strconv"
)

func main() {

	v := "3.1415926535"
	v1, _ := strconv.ParseFloat(v, 32)
	v2, _ := strconv.ParseFloat(v, 64)
	value, _ := strconv.ParseFloat(fmt.Sprintf("%.2f", v1), 64)

	fmt.Println(value)
	fmt.Println(v2)

	var data = []int{-1, 0, 1, 2, -1, -4}
	//var twoData = [][]int{{0, 1, 2, 3}}
	//fmt.Println(twoData[0][1])
	//
	//sort.MergeSort(data, len(data))

	fmt.Println(data)
	res := letcode.ThreeSum(data)
	_ = res
	fmt.Println(res)
}
