package main

import (
	"fmt"
)

func main() {

	//fmt.Println(leetcode.SearchMatrix([][]int{
	//	{-5},
	//}, -10))

	fmt.Println(test(10))

}

func test(n int) [][]int {
	var sil = make([][]int, n)

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			sil[i] = append(sil[i], j)
		}
	}
	return sil
}

func sum(n, m int) int {
	sum := 0
	for i := 0; i < n; i++ {
		if m == n {
			return sum
		}
		sum += i * 2
	}
	return sum
}
