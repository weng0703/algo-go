package linkedlist

import "testing"

var l *LinkedList

func init() {
	n5 := &ListNode{value: 5}
	n4 := &ListNode{value: 4, Next: n5}
	n3 := &ListNode{value: 3, Next: n4}
	n2 := &ListNode{value: 2, Next: n3}
	n1 := &ListNode{value: 1, Next: n2}
	l = &LinkedList{head: &ListNode{Next: n1}}
}

func TestInsertToHead(t *testing.T) {
	l := NewlinkedList()
	for i := 0; i < 10; i++ {
		l.InsertToHead(i + 1)
	}
	l.Print()
}
func TestInsertToTail(t *testing.T) {
	l := NewlinkedList()
	for i := 0; i < 10; i++ {
		l.InsertToTail(i + 1)
	}
	l.Print()
}

func TestFindByIndex(t *testing.T) {
	l := NewlinkedList()
	for i := 0; i < 10; i++ {
		l.InsertToTail(i + 1)
	}
	t.Log(l.FindByIndex(0))
	t.Log(l.FindByIndex(9))
	t.Log(l.FindByIndex(5))
	t.Log(l.FindByIndex(11))
}

func TestDeleteNode(t *testing.T) {
	l := NewlinkedList()
	for i := 0; i < 3; i++ {
		l.InsertToTail(i + 1)
	}
	l.Print()
	t.Log(l.DeleteNode(l.head.Next))
	l.Print()

	t.Log(l.DeleteNode(l.head.Next.Next))
	l.Print()
}

func TestPalindrome1(t *testing.T) {
	strs := []string{"heooeh", "hello", "heoeh", "a", ""}
	for _, str1 := range strs {
		l := NewlinkedList()
		for _, c := range str1 {
			l.InsertToTail(string(c))
		}
		l.Print()
		t.Log(IsPalindrome1(l))

	}
}

func TestMergeSortedList(t *testing.T) {
	n5 := &ListNode{value: 9}
	n4 := &ListNode{value: 7, Next: n5}
	n3 := &ListNode{value: 5, Next: n4}
	n2 := &ListNode{value: 3, Next: n3}
	n1 := &ListNode{value: 1, Next: n2}
	l1 := &LinkedList{head: &ListNode{Next: n1}}

	n10 := &ListNode{value: 10}
	n9 := &ListNode{value: 8, Next: n10}
	n8 := &ListNode{value: 6, Next: n9}
	n7 := &ListNode{value: 4, Next: n8}
	n6 := &ListNode{value: 2, Next: n7}
	l2 := &LinkedList{head: &ListNode{Next: n6}}

	MergeSortedList(l1, l2).Print()
}
func TestDeleteBottomN(t *testing.T) {
	l.Print()
	l.DeleteBottomN(2)
	l.Print()
}

func TestLinkedList_Reverse(t *testing.T) {
	l.Print()
	l.Reverse()
	l.Print()
}
