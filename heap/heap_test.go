package heap

import (
	"fmt"
	"testing"
)

func TestHeap_Insert(t *testing.T) {
	heap := NewHeap(5)
	heap.Insert(1)
	heap.Insert(2)
	heap.Insert(3)
	heap.Insert(5)
	fmt.Println(heap.arr)
}
