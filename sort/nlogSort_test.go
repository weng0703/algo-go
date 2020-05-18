package sort

import "testing"

func TestMergeSort(t *testing.T) {
	arr := []int{5, 4}
	MergeSort(arr)
	t.Log(arr)

	arr = []int{5, 3, 4, 6, 1}
	MergeSort(arr)
	t.Log(arr)
}
