package custom_list

import "strconv"

type ListIterator struct {
}

func (*ListIterator) Inc() {

}

type ListNode struct {
	Value int
	Next  *ListNode
	Prev  *ListNode
}

type List struct {
	Size uint32
	Head ListNode
}

func NewList() List {
	var l List
	l.Head.Next = &l.Head
	l.Head.Prev = &l.Head
	return l
}

func (list *List) PushBack(val int) {
	list.Head.Prev = &ListNode{Value: val, Next: &list.Head, Prev: list.Head.Prev}
}

func (list *List) PushFront(val int) {
	list.Head.Next = &ListNode{Value: val, Next: list.Head.Next, Prev: list.Head.Prev}
}

func (list *List) String() string {
	var result string
	for it := list.Head.Next; it != &list.Head; it = it.Next {
		result = result + strconv.FormatInt(int64(it.Value), 10)
	}
	return result
}

func foo() {
	var l List = NewList()
	l.PushBack(1)
}
