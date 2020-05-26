package heap

type Heap struct {
	arr   []int
	n     int
	count int
}

func NewHeap(capacity int) *Heap {
	return &Heap{make([]int, capacity+1), capacity, 0}
}

func (heap *Heap) Insert(data int) {
	if heap.count == heap.n {
		return
	}
	heap.count++
	heap.arr[heap.count] = data
	i := heap.count
	parent := i / 2
	//堆化
	for parent > 0 && heap.arr[parent] < heap.arr[i] {
		heap.arr[parent], heap.arr[i] = heap.arr[i], heap.arr[parent]
		i = parent
		parent = i / 2
	}
}

func (heap *Heap) removeMax() {
	if heap.count == 0 {
		return
	}
	heap.arr[heap.count], heap.arr[1] = heap.arr[1], heap.arr[heap.count]
	heap.count--
	heapifyUpToDown(heap.arr, heap.count)
}

func heapifyUpToDown(arr []int, count int) {
	for i := 1; i <= count/2; {
		maxIndex := i
		if arr[i] < arr[i*2] {
			maxIndex = i * 2
		}
		if i*2+1 <= count && arr[maxIndex] < arr[i*2+1] {
			maxIndex = i*2 + 1
		}
		if maxIndex == i { // 未交换
			break
		}
		arr[maxIndex], arr[i] = arr[i], arr[maxIndex]
		i = maxIndex
	}
}

//建堆
func buildHeap(arr []int, n int) {
	for i := n / 2; i >= 1; i-- {
		heapifyUpToDown1(arr, i, n)
	}
}

func heapifyUpToDown1(arr []int, top int, count int) {
	for i := top; i <= count/2; { //局部堆化，以当前节点为顶部节点
		maxIndex := i
		if arr[i] < arr[i*2] {
			maxIndex = i * 2
		}
		if i*2+1 <= count && arr[maxIndex] < arr[i*2+1] {
			maxIndex = i*2 + 1
		}
		if maxIndex == i { // 未交换
			break
		}
		arr[maxIndex], arr[i] = arr[i], arr[maxIndex]
		i = maxIndex
	}
}

//堆排序
func sort(arr []int, n int) {
	buildHeap(arr, n)
	k := n
	for k >= 1 {
		arr[1], arr[k] = arr[k], arr[1]
		heapifyUpToDown1(arr, 1, k-1)
		k--
	}
}
