package custom_list

import (
	"errors"
	"fmt"
)

type ListIterator struct {
}

func (*ListIterator) Inc() {

}

type ListNode struct {
	Value interface{}
	Next  *ListNode
	Prev  *ListNode
}

type List interface {
	PushBack(interface{})
	PushFront(interface{})
	Find(interface{}) (interface{}, error)
	String() string
}

type listImpl struct {
	Size uint32
	Head ListNode
}

func NewList() List {
	l := new(listImpl)
	l.Head.Next = &l.Head
	l.Head.Prev = &l.Head
	return l
}

func (list *listImpl) PushBack(val interface{}) {
	list.Head.Prev = &ListNode{Value: val, Next: &list.Head, Prev: list.Head.Prev}
}

func (list *listImpl) PushFront(val interface{}) {
	list.Head.Next = &ListNode{Value: val, Next: list.Head.Next, Prev: list.Head.Prev}
}

func (list *listImpl) Find(val interface{}) (interface{}, error) {
	for it := list.Head.Next; it != &list.Head; it = it.Next {
		if it.Value == val {
			return it.Value, nil
		}
	}
	return nil, errors.New("Not found")
}

func (list *listImpl) Begin()

func (list *listImpl) End()

func (list *listImpl) Insert()

func (list *listImpl) String() string {
	var result string
	for it := list.Head.Next; it != &list.Head; it = it.Next {
		if it != list.Head.Next {
			result += " "
		}
		result += fmt.Sprintf("%v", it.Value)
	}
	return result
}
