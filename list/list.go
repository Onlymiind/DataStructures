package list

import (
	"fmt"

	"github.com/Onlymiind/DataStructures/iterator"
)

type listNode struct {
	Value interface{}
	Next  *listNode
	Prev  *listNode
}

type listIterator struct {
	Node *listNode
}

func (it *listIterator) Inc() {
	it.Node = it.Node.Next
}

func (it *listIterator) Dec() {
	it.Node = it.Node.Prev
}

func (it *listIterator) Get() interface{} {
	return it.Node.Value
}

func (it *listIterator) Set(val interface{}) {
	it.Node.Value = val
}

func (it listIterator) Equal(other iterator.Iterator) bool {
	return it.Node == other.(*listIterator).Node
}



type List interface {
	PushBack(interface{})
	PushFront(interface{})
	Find(interface{}) iterator.Iterator
	Insert(iterator.Iterator, interface{})
	Erase(iterator.Iterator)
	Clear()
	Size() uint64
	Empty() bool
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

func (l *listImpl) PushBack(val interface{}) {
	temp := &listNode{Value: val, Next: &l.Head, Prev: l.Head.Prev}
	l.Head.Prev.Next = temp
	l.Head.Prev = temp
	l.ElemCount++
}

func (l *listImpl) PushFront(val interface{}) {
	temp := &listNode{Value: val, Next: l.Head.Next, Prev: &l.Head}
	l.Head.Next.Prev = temp
	l.Head.Next = temp
	l.ElemCount++
}

func (l *listImpl) Find(val interface{}) iterator.Iterator {
	for it := l.Begin(); !it.Equal(l.End()); it.Inc() {
		if it.Get() == val {
			return it
		}
	}
	return l.End()
}

func (l *listImpl) Begin() iterator.Iterator {
	return &listIterator{l.Head.Next}
}

func (l *listImpl) End() iterator.Iterator {
	return &listIterator{&l.Head}
}

func (l *listImpl) Insert(before iterator.Iterator, val interface{}) {
	pos := before.(*listIterator).Node
	temp := &listNode{Value: val, Next : pos, Prev : pos.Prev}
	pos.Prev.Next = temp
	pos.Prev = temp
	l.ElemCount++
}

func (l *listImpl) Erase(at iterator.Iterator) {
	pos := at.(*listIterator).Node
	pos.Prev.Next = pos.Next
	pos.Next.Prev = pos.Prev
	if l.ElemCount != 0 {
		l.ElemCount--
	}
}

func (l *listImpl) Clear() {
	l.Head.Next = &l.Head
	l.Head.Prev = &l.Head
	l.ElemCount = 0
}

func (l *listImpl) Size() uint64 {
	return l.ElemCount
}

func (l *listImpl) Empty() bool {
	return l.Size() == 0
}

func (l *listImpl) String() string {
	var result string
	for it := l.Begin(); !it.Equal(l.End()); it.Inc() {
		if !it.Equal(l.Begin()) {
			result += " "
		}
		result += fmt.Sprintf("%v", it.Get())
	}
	return result
}
