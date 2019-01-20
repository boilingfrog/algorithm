package main

import (
	"algorithm/letcode"
	"fmt"
)

func main() {
	var arr = []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
	res := letcode.MaxArea(arr)
	fmt.Println(res)
}
