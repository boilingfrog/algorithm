package main

import (
	"fmt"
)

var arr = [100]int{}

//数组指针
var data *[100]int = &arr

var count = 0

func main() {
	//Insert(2)
	//Insert(3)
	//Insert(5)
	//Insert(7)
	//Insert(89)
	//Insert(9)
	//Insert(8)
	//Insert(4)

	var arr = []int{2, 3, 4, 5, 6, 7, 8, 9, 4}

	MaxHeap(arr, len(arr))

	fmt.Println(data)
}

// MaxHeap 将一个数组转换成大顶堆
func MaxHeap(arrData []int, n int) {
	//var data *[100]int = &arr

	for i := 0; i < n; i++ {
		data[i+1] = arrData[i]
	}
	count = n
	for i := count / 2; i >= 1; i-- {
		ShiftDown(i)
	}

}

// ExtractMax 从堆中取出一个值
func ExtractMax() int {
	if count <= 0 {
		return 0
	}
	ret := data[1]
	SwapArr(1, count)
	count--
	ShiftDown(1)
	return ret
}

// ShiftDown ...
func ShiftDown(k int) {
	for {
		if 2*k > count {
			break
		}
		j := 2 * k
		if j+1 <= count && data[j+1] > data[j] {
			j += 1
		}
		if data[k] >= data[j] {
			break
		}
		SwapArr(j, k)
		k = j
	}
}

// Insert ...
func Insert(item int) {
	data[count+1] = item
	count++
	shiftUp(count)
}

// shiftUp ...
func shiftUp(k int) {
	for {
		if k <= 1 || data[k/2] >= data[k] {
			break
		}
		SwapArr(k/2, k)
		k = k / 2
	}
}

// SwapArr ...
func SwapArr(j int, i int) {
	t := data[j]
	arr[j] = arr[i]
	arr[i] = t
}
