package sort

import (
	"math/rand"
	"testing"
)

func TestMergeSort(t *testing.T) {
	arr := []int{5, 4}
	MergeSort(arr)
	t.Log(arr)

	arr = []int{5, 3, 4, 6, 1}
	MergeSort(arr)
	t.Log(arr)
}

func createRandomArr(length int) []int {
	arr := make([]int, length, length)
	for i := 0; i < length; i++ {
		arr[i] = rand.Intn(100)
	}
	return arr
}
func TestQuickSort(t *testing.T) {
	arr := []int{5, 4}
	QuickSort(arr)
	t.Log(arr)

	arr = createRandomArr(100)
	QuickSort(arr)
	t.Log(arr)
}
