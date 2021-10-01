package main

func bucketSort(arr []int) {
	if len(arr) <= 1 {
		return
	}

	max, min := arr[0], arr[0]
	for _, item := range arr {
		if item < min {
			min = item
		} else if item > max {
			max = item
		}
	}

	bucketCount := make([][]int, (max-min)%10+1)
	for _, item := range arr {
		bucketCount[(item-min)%10] = append(bucketCount[(item-min)%10], item)
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
