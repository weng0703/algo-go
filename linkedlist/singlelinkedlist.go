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
