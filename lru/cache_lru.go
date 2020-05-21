package lru

const (
	HOST_BIT = uint64(^uint(0)) == ^uint64(0)
	LENGTH   = 100
)

type lruNode struct {
	prev  *lruNode
	next  *lruNode
	key   int
	value int
	hnext *lruNode // hash拉链
}

type LRUCache struct {
	node     []lruNode
	head     *lruNode //头结点
	tail     *lruNode //尾结点
	capacity int
	used     int
}

func Constructor(capacity int) LRUCache {
	return LRUCache{
		node:     make([]lruNode, LENGTH),
		head:     nil,
		tail:     nil,
		capacity: capacity,
		used:     0,
	}
}

func hash(key int) int {
	if HOST_BIT { //64位系统
		return (key ^ (key >> 32)) & (LENGTH - 1)
	}
	return (key ^ (key >> 16)) & (LENGTH - 1) // 32位系统
}
func (this *LRUCache) addNode(key int, value int) {
	newNode := &lruNode{
		key:   key,
		value: value,
	}
	tmp := &this.node[hash(key)]
	newNode.hnext = tmp.hnext
	tmp.hnext = newNode
	this.used++
	if this.tail == nil {
		this.tail, this.head = newNode, newNode
		return
	}
	//插入到队尾
	this.tail.next = newNode
	newNode.prev = this.tail
	this.tail = newNode
}

func (this *LRUCache) delNode() {
	if this.head == nil {
		return
	}
	prev := &this.node[hash(this.head.key)]
	tmp := prev.hnext
	for tmp != nil && tmp.key != this.head.key {
		prev = tmp
		tmp = tmp.hnext
	}
	if tmp == nil { //没有找到对应的结点
		return
	}
	prev.hnext = tmp.hnext
	this.head = this.head.next
	this.head.prev = nil
	this.used--
}

func (this *LRUCache) searchNode(key int) *lruNode {
	if this.tail == nil {
		return nil
	}
	tmp := this.node[hash(key)].hnext
	for tmp != nil {
		if tmp.key == key {
			return tmp
		}
		tmp = tmp.hnext
	}
	return nil
}

func (this *LRUCache) moveToTail(node *lruNode) {
	if this.tail == node {
		return
	}
	if this.head == node {
		this.head = node.next
		this.head.prev = nil
	} else {
		node.next.prev = node.prev
		node.prev.next = node.next
	}
	node.next = nil
	this.tail.next = node
	node.prev = this.tail
	this.tail = node
}

func (this *LRUCache) Get(key int) int {
	if this.tail == nil {
		return -1
	}
	if tmp := this.searchNode(key); tmp != nil {
		this.moveToTail(tmp)
		return tmp.key
	}
	return -1
}
func (this *LRUCache) Put(key int, value int) {
	if tmp := this.searchNode(key); tmp != nil {
		tmp.value = value
		this.moveToTail(tmp)
		return
	}
	this.addNode(key, value)
	if this.used > this.capacity {
		this.delNode()
	}
}
