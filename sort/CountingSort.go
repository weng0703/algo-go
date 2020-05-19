package sort

func CountingSort(a []int, n int) {
	if n <= 1 {
		return
	}
	max := getMax(a)
	c := make([]int, max+1)
	// 计数
	for i := range a {
		c[a[i]]++
	}
	//让当前索引指定数累计为大于等于当前索引的个数和
	for i := 1; i <= max; i++ {
		c[i] += c[i-1]
	}
	r := make([]int, n)
	for i := n - 1; i >= 0; i-- {
		index := c[a[i]] - 1
		r[index] = a[i]
		c[a[i]]--
	}
	copy(a, r)
}
