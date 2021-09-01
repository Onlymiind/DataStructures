package queue

import (
	"errors"
)

type Queue interface {
	Enqueue(interface{}) error
	Dequeue() error
	Peek() interface{}
	Empty() bool
	Size() uint64
	Capacity() uint64
	Clear()
}

type intQueue struct {
	Elements  []int
	ElemCount uint64
	Head      uint64
	Tail      uint64
}

func NewIntQueue(size uint64) Queue {
	return &intQueue{make([]int, size), 0, 0, 0}
}

func (q *intQueue) Enqueue(val interface{}) error {
	value := val.(int)
	if q.Size() == q.Capacity() {
		return errors.New("Queue capacity exhausted")
	}
	q.Elements[q.Head] = value
	q.ElemCount++
	q.Head = (q.Head + 1) % q.Capacity()
	return nil
}

func (q *intQueue) Dequeue() error {
	if q.Empty() {
		return errors.New("Trying to remove an element from empty queue")
	}

	q.Tail = (q.Tail + 1) % q.Capacity()
	q.ElemCount--
	return nil
}

func (q *intQueue) Peek() interface{} {
	return q.Elements[q.Tail]
}

func (q *intQueue) Size() uint64 {
	return q.ElemCount
}

func (q *intQueue) Capacity() uint64 {
	return uint64(len(q.Elements))
}

func (q *intQueue) Empty() bool {
	return q.ElemCount == 0
}

func (q *intQueue) Clear() {
	q.ElemCount, q.Head, q.Tail = 0, 0, 0
}
