package main

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

// Max golang max int
func Max(first int, args ...int) int {
	for _, v := range args {
		if first < v {
			first = v
		}
	}
	return first
}

// Min golang min int
func Min(first int, args ...int) int {
	for _, v := range args {
		if first > v {
			first = v
		}
	}
	return first
}

func Swap(arr []int, j int, i int) {
	t := arr[j]
	arr[j] = arr[i]
	arr[i] = t
}
