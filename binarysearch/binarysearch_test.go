package binarysearch

import "testing"

func TestBinarySearch(t *testing.T) {
	var a []int

	a = []int{1, 3, 5, 6, 8}
	search := BinarySearch(a, 31)
	t.Log(search)
}

func TestBinarySearchRecursive(t *testing.T) {
	var a []int

	a = []int{1, 3, 5, 6, 8}
	search := BinarySearchRecursive(a, 31)
	t.Log(search)
}

func TestBinarySearchFirst(t *testing.T) {
	var a []int

	a = []int{1, 2, 2, 2, 3, 4}
	index := BinarySearchFirst(a, 2)
	index2 := BinarySearch(a, 2)
	t.Log(index, index2)
}

func TestBinarySearchLast(t *testing.T) {
	var a []int

	a = []int{1, 2, 2, 2, 3, 4}
	index := BinarySearchLast(a, 2)
	index2 := BinarySearch(a, 2)
	t.Log(index, index2)
}
