package list

import (
	"fmt"

	"github.com/Onlymiind/DataStructures/iterator"
)

type listIterator listNode

func (it *listIterator) Inc() {
	listNode(it).Node = it.Node.Next
}

func (it *listIterator) Dec() {
	listNode(it).Node = it.Node.Prev
}

func (it *listIterator) Get() interface{} {
	return it.Node.Value
}

func (it *listIterator) Set(val interface{}) {
	listNode(it).Node.Value = val
}

func (it *listIterator) Equal(other iterator.Iterator) bool {
	other_it := other.(*listIterator)
	return listNode(it).Node == listNode(other_it).Node
}

type listNode struct {
	Value interface{}
	Next  *listNode
	Prev  *listNode
}

type List interface {
	PushBack(interface{})
	PushFront(interface{})
	Find(interface{}) iterator.Iterator
	Insert(iterator.Iterator, interface{})
	Erase(iterator.Iterator)
	Clear()
	Size() uint64
	String() string
	Begin() iterator.Iterator
	End() iterator.Iterator
}

type listImpl struct {
	ElemCount uint64
	Head listNode
}

func NewList() List {
	l := new(listImpl)
	l.Head.Next = &l.Head
	l.Head.Prev = &l.Head
	return l
}

func (listImpl *listImpl) PushBack(val interface{}) {
	temp := &listNode{Value: val, Next: &listImpl.Head, Prev: listImpl.Head.Prev}
	listImpl.Head.Prev.Next = temp
	listImpl.Head.Prev = temp
	listImpl.ElemCount++
}

func (listImpl *listImpl) PushFront(val interface{}) {
	temp := &listNode{Value: val, Next: listImpl.Head.Next, Prev: &listImpl.Head}
	listImpl.Head.Next.Prev = temp
	listImpl.Head.Next = temp
	listImpl.ElemCount++
}

func (listImpl *listImpl) Find(val interface{}) iterator.Iterator {
	for it := listImpl.Begin(); !it.Equal(listImpl.End()); it.Inc() {
		if it.Get() == val {
			return it
		}
	}
	return listImpl.End()
}

func (listImpl *listImpl) Begin() iterator.Iterator {
	return &listIterator{listImpl.Head.Next}
}

func (listImpl *listImpl) End() iterator.Iterator {
	return &listIterator{&listImpl.Head}
}

func (listImpl *listImpl) Insert(before iterator.Iterator, val interface{}) {
	pos := before.(*listIterator)
	temp := &listNode{Value: val, Next : pos.Node, Prev : pos.Node.Prev}
	pos.Node.Prev.Next = temp
}

func (listImpl *listImpl) Erase(at iterator.Iterator) {
	pos := at.(*listIterator)
	pos.Node.Prev.Next = pos.Node.Next
	pos.Node.Next.Prev = pos.Node.Prev
}

func (listImpl *listImpl) Clear() {
	listImpl.Head.Next = &listImpl.Head
	listImpl.Head.Prev = &listImpl.Head
	listImpl.ElemCount = 0
}

func (listImpl *listImpl) Size() uint64 {
	return listImpl.ElemCount
}

func (listImpl *listImpl) String() string {
	var result string
	for it := listImpl.Begin(); !it.Equal(listImpl.End()); it.Inc() {
		if it != listImpl.Begin() {
			result += " "
		}
		result += fmt.Sprintf("%v", it.Get())
	}
	return result
}
