package main

func countingSort(arr []int) {
	if len(arr) <= 1 {
		return
	}

	max := arr[0]
	for _, item := range arr {
		if item > max {
			max = item
		}
	}

	bucketCount := make([][]int, max+1)
	for _, item := range arr {
		bucketCount[item] = append(bucketCount[item], item)
	}

	i := 0
	for _, bucketArr := range bucketCount {
		if len(bucketArr) > 1 {
			QuickSort3Ways(bucketArr, len(bucketArr)-1)
		}
		for _, item := range bucketArr {
			arr[i] = item
			i++
		}
	}
}
