package main

import "fmt"

func main() {

	//fmt.Println(letcode.SearchMatrix([][]int{
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
