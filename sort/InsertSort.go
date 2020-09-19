package main

import "fmt"

func insert(arr []int) {
	for i := 1; i < len(arr); i++ {
		key := i - 1
		tmp := arr[i]
		for key >= 0 && tmp < arr[key] {
			arr[key+1] = arr[key]
			key--
		}
		if key+1 != i {
			arr[key+1] = tmp
		}
	}
	fmt.Println(arr)
}
