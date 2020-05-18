package sort

import (
	"fmt"
	"testing"
)

func TestBubbleSort(t *testing.T) {
	var a = []int64{1, 23, 46, 788, 5, 6, 44}
	BubbleSort(a, len(a))
	fmt.Println(a)
}

func TestInsertionSort(t *testing.T) {
	var a = []int64{24, 23, 46, 788, 5, 6, 44}
	InsertionSort(a, len(a))
	fmt.Println(a)
}

func TestSelectionSort(t *testing.T) {
	var a = []int64{24, 23, 46, 788, 5, 6, 44}
	InsertionSort(a, len(a))
	fmt.Println(a)
}
