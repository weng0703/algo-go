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

//快速排序
func QuickSort(arr []int) {
	separateSort(arr, 0, len(arr)-1)
}

func separateSort(arr []int, start, end int) {
	if start >= end {
		return
	}
	i := partition(arr, start, end)
	separateSort(arr, start, i-1)
	separateSort(arr, i+1, end)
}

//获取分区点
func partition(arr []int, start, end int) int {
	pivot := end //选取最后一位当做对比数字
	i := start
	for j := start; j < end; j++ {
		if arr[j] < arr[pivot] {
			if !(i == j) {
				arr[j], arr[i] = arr[i], arr[j]
			}
			i++
		}
	}
	arr[i], arr[end] = arr[end], arr[i]
	return i
}
