package sort

//归并排序
func MergeSort(arr []int) {
	arrLen := len(arr)
	if arrLen <= 1 {
		return
	}
	mergeSort(arr, 0, arrLen-1)
}

//归
func mergeSort(arr []int, start, end int) {
	if start >= end {
		return
	}
	mid := (start + end) / 2
	mergeSort(arr, start, mid)
	mergeSort(arr, mid+1, end)
	merge(arr, start, mid, end)
}

//合并
func merge(arr []int, start, mid, end int) {
	tempArr := make([]int, end-start+1)
	i := start
	j := mid + 1
	k := 0
	//将两个有序数组按序装入临时数组
	for ; i <= mid && j <= end; k++ {
		if arr[i] > arr[j] {
			tempArr[k] = arr[i]
			i++
		} else {
			tempArr[k] = arr[j]
			j++
		}
	}
	for ; i <= mid; i++ {
		tempArr[k] = arr[i]
		k++
	}
	for ; j <= end; j++ {
		tempArr[k] = arr[j]
		k++
	}
	copy(arr[start:end+1], tempArr)
}
