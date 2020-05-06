package stack

import "fmt"

type node struct {
	next  *node
	value interface{}
}

type LinkedListStack struct {
	TopNode *node
}

func NewLinkedListStack() *LinkedListStack {
	return &LinkedListStack{nil}
}
func (this *LinkedListStack) IsEmpty() bool {
	if this.TopNode == nil {
		return true
	}
	return false
}
func (this *LinkedListStack) Push(v interface{}) {
	this.TopNode = &node{this.TopNode, v}
}
func (this *LinkedListStack) Pop() interface{} {
	if this.IsEmpty() {
		return nil
	}
	v := this.TopNode.value
	this.TopNode = this.TopNode.next
	return v
}
func (this *LinkedListStack) Top() interface{} {
	if this.IsEmpty() {
		return nil
	}

	return this.TopNode.value
}

func (this *LinkedListStack) Flush() {
	this.TopNode = nil
}

func (this *LinkedListStack) Print() {
	if this.IsEmpty() {
		fmt.Println("empty stack")
	} else {
		cur := this.TopNode
		for nil != cur {
			fmt.Println(cur.value)
			cur = cur.next
		}
	}
}
