package list

import (
	"errors"
	"fmt"

	"github.com/Onlymiind/DataStructures/iterator"
)

type intIterator struct {
	Node *ListNode
}

func (it *intIterator) Inc() {
	it.Node = it.Node.Next
}

func (it *intIterator) Dec() {
	it.Node = it.Node.Prev
}

func (it *intIterator) Get() interface{} {
	return it.Node.Value
}

func (it *intIterator) Set(val interface{}) {
	value := val.(int)
	it.Node.Value = value
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
	Insert(iterator.Iterator, interface{})
	Erase(iterator.Iterator)
	Clear()
	Size() uint64
	String() string
	Begin() iterator.Iterator
	End() iterator.Iterator
}

type intList struct {
	ElemCount uint64
	Head ListNode
}

func NewIntList() List {
	l := new(intList)
	l.Head.Next = &l.Head
	l.Head.Prev = &l.Head
	return l
}

func (list *intList) PushBack(val interface{}) {
	value := val.(int)
	temp := &ListNode{Value: value, Next: &list.Head, Prev: list.Head.Prev}
	list.Head.Prev.Next = temp
	list.Head.Prev = temp
	if(list.ElemCount == 0) {
		list.Head.Next = list.Head.Prev
	}
	list.ElemCount++
}

func (list *intList) PushFront(val interface{}) {
	value := val.(int)
	temp := &ListNode{Value: value, Next: list.Head.Next, Prev: list.Head.Prev}
	list.Head.Next.Prev = temp
	list.Head.Next = temp
	if(list.ElemCount == 0) {
		list.Head.Prev = list.Head.Next
	}
	list.ElemCount++
}

func (list *intList) Find(val interface{}) (interface{}, error) {
	for it := list.Head.Next; it != &list.Head; it = it.Next {
		if it.Value == val {
			return it.Value, nil
		}
	}
	return nil, errors.New("Not found")
}

func (list *intList) Begin() iterator.Iterator {
	return &intIterator{list.Head.Next}
}

func (list *intList) End() iterator.Iterator {
	return &intIterator{&list.Head}
}

func (list *intList) Insert(before iterator.Iterator, val interface{}) {
	pos := before.(*intIterator)
	value := val.(int)
	temp := &ListNode{Value: value, Next : pos.Node, Prev : pos.Node.Prev}
	pos.Node.Prev.Next = temp
}

func (list *intList) Erase(at iterator.Iterator) {
	pos := at.(*intIterator)
	pos.Node.Prev.Next = pos.Node.Next
	pos.Node.Next.Prev = pos.Node.Prev
}

func (list *intList) Clear() {
	list.Head.Next = &list.Head
	list.Head.Prev = &list.Head
	list.ElemCount = 0
}

func (list *intList) Size() uint64 {
	return list.ElemCount
}

func (list *intList) String() string {
	var result string
	for it := list.Begin(); it != list.End(); it.Inc() {
		if it != list.Begin() {
			result += " "
		}
		result += fmt.Sprintf("%v", it.Get())
	}
	return result
}
