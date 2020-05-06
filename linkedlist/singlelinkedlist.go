package linkedlist

import "fmt"

//节点
type ListNode struct {
	Next  *ListNode
	value interface{}
}

type LinkedList struct {
	head   *ListNode //头结点
	length uint      //链表长度
}

func NewListNode(v interface{}) *ListNode {
	return &ListNode{
		nil,
		v,
	}
}

func (this *ListNode) GetValue() interface{} {
	return this.value
}

func (this *ListNode) GetNext() *ListNode {
	return this.Next
}

func NewlinkedList() *LinkedList {
	return &LinkedList{
		NewListNode(nil),
		0,
	}
}

//再某个节点后插入
func (this *LinkedList) InsertAfter(p *ListNode, v interface{}) bool {
	if p == nil {
		return false
	}
	newNode := NewListNode(v)
	oldNode := p.Next
	p.Next = newNode
	newNode.Next = oldNode
	this.length++
	return true
}

//在某个节点前面插入
func (this *LinkedList) InsertBefor(p *ListNode, v interface{}) bool {
	if p == nil || p == this.head {
		return false
	}
	cur := this.head.Next
	pre := this.head
	for cur != nil {
		if cur == p {
			break
		}
		pre = cur
		cur = cur.Next
	}
	if cur == nil { //没有找到对应的节点
		return false
	}
	newNode := NewListNode(v)
	pre.Next = newNode
	newNode.Next = cur
	this.length++
	return true
}

//头插
func (this *LinkedList) InsertToHead(v interface{}) bool {
	return this.InsertAfter(this.head, v)
}

//尾插
func (this *LinkedList) InsertToTail(v interface{}) bool {
	cur := this.head
	for cur.Next != nil {
		cur = cur.Next
	}
	return this.InsertAfter(cur, v)
}

//通过索引查找节点
func (this *LinkedList) FindByIndex(index uint) *ListNode {
	if index >= this.length { //索引位置超出
		return nil
	}
	cur := this.head.Next
	var i uint = 0
	for ; i < index; i++ {
		cur = cur.Next
	}
	return cur
}

//删除制定节点
func (this *LinkedList) DeleteNode(p *ListNode) bool {
	if p == nil {
		return false
	}
	cur := this.head.Next
	pre := this.head
	for cur != nil {
		if cur == p {
			break
		}
		pre = cur
		cur = cur.Next
	}
	if cur == nil { //未找到制定节点
		return false
	}
	pre.Next = cur.Next
	cur = nil
	this.length--
	return true
}

//打印链表
func (this *LinkedList) Print() {
	cur := this.head.Next
	format := ""
	for cur != nil {
		format += fmt.Sprintf("%+v", cur.GetValue())
		cur = cur.Next
		if nil != cur {
			format += "->"
		}
	}
	fmt.Println(format)
}

//判断是否回文链表
//用数组暂存前半段
func IsPalindrome1(l *LinkedList) bool {
	lLen := l.length
	if lLen == 0 {
		return false
	}
	if lLen == 1 {
		return true
	}
	s := make([]string, 0, lLen/2)
	cur := l.head
	var i uint = 1
	for cur.Next != nil {
		cur = cur.Next
		if lLen%2 != 0 && i == lLen/2+1 { //链表数为奇数个，跳过中间点
			i++
			continue
		}
		if i <= lLen/2 {
			s = append(s, cur.GetValue().(string))
		} else {
			if s[lLen-i] != cur.GetValue().(string) {
				return false
			}
		}
		i++
	}
	return true
}

//合并有序链表
func MergeSortedList(l1, l2 *LinkedList) *LinkedList {
	if nil == l1 || nil == l1.head || nil == l1.head.Next {
		return l2
	}
	if nil == l2 || nil == l2.head || nil == l2.head.Next {
		return l1
	}
	l := NewlinkedList()
	cur := l.head
	cur1 := l1.head.Next
	cur2 := l2.head.Next
	for cur1 != nil && cur2 != nil {
		if cur1.value.(int) > cur2.value.(int) {
			cur.Next = cur2
			cur2 = cur2.Next
		} else {
			cur.Next = cur1
			cur1 = cur1.Next
		}
		cur = cur.Next
	}
	if nil != cur1 {
		cur.Next = cur1
	} else if cur2 != nil {
		cur.Next = cur2
	}
	return l
}

//删除倒数第n个节点
func (this *LinkedList) DeleteBottomN(n uint) bool {
	if nil == this.head || nil == this.head.Next {
		return false
	}
	fast := this.head
	for i := uint(1); i <= n && fast != nil; i++ {
		fast = fast.Next
	}
	if nil == fast {
		return false
	}
	slow := this.head
	for nil != fast.Next {
		slow = slow.Next
		fast = fast.Next
	}
	slow.Next = slow.Next.Next
	return true
}

//单链表反转
func (this *LinkedList) Reverse() {
	if nil == this.head || nil == this.head.Next || nil == this.head.Next.Next {
		return
	}
	var pre *ListNode = nil
	cur := this.head.Next
	for nil != cur {
		tem := cur.Next
		cur.Next = pre
		pre = cur
		cur = tem
	}
	this.head.Next = pre
}
