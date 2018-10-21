package sort

// 归并排序(自低向上)
func MergeSortUp(arr []int, n int) {

	for sz := 1; sz <= n; sz += sz {

		for i := 0; i+sz < n; i += sz + sz {
			Merge(arr,i,i+sz-1,Max(i+sz+sz-1,n-1))
		}

	}

}


