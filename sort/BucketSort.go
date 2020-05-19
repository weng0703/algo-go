package sort

import "fmt"

//获取待排序数组中的最大值
func getMax(a []int) int {
	var max = a[0]
	for i := 1; i < len(a); i++ {
		if a[i] > max {
			max = a[i]
		}
	}
	return max
}

func BucketSort(a []int) {
	num := len(a)
	if num <= 1 {
		return
	}
	max := getMax(a)
	buckets := make([][]int, num)
	index := 0
	for i := 0; i < num; i++ {
		index = a[i] * (num - 1) / max                //算出桶序号
		buckets[index] = append(buckets[index], a[i]) //加入对应的桶
	}
	tmpPos := 0
	for i := 0; i < num; i++ {
		bucketLen := len(buckets[i])
		if bucketLen > 0 {
			QuickSort(buckets[i]) // 桶内做快速排序
			copy(a[tmpPos:], buckets[i])
			tmpPos += bucketLen
		}
	}
}

//简单实现--特殊情况简化为计数排序
func BucketSortSimple(a []int) {
	if len(a) <= 1 {
		return
	}
	// 分入桶里
	array := make([]int, getMax(a)+1)
	for i := 0; i < len(a); i++ {
		array[a[i]]++
	}
	fmt.Println(array)
	c := make([]int, 0)
	for i := 0; i < len(array); i++ {
		for array[i] != 0 {
			c = append(c, i)
			array[i]--
		}
	}
	copy(a, c)
}
