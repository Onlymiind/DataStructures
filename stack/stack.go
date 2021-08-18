package stack

import (
	"errors"
)

type Stack interface {
	Pop()
	Push(interface{}) error
	Top() interface{}
	Size() uint64
	Capacity() uint64
}

type intStack struct {
	Elements  []int
	ElemCount uint64
}

func NewIntStack(size uint64) Stack {
	var s intStack
	s.Elements = make([]int, size)
	return &s
}

func (s *intStack) Pop() {
	if s.ElemCount != 0 {
		s.ElemCount--
	}
}

func (s *intStack) Push(val interface{}) error {
	if s.ElemCount == uint64(len(s.Elements)) {
		return errors.New("Not enough space")
	}
	value := val.(int)
	s.Elements[s.ElemCount] = value
	s.ElemCount++
	return nil
}

func (s *intStack) Top() interface{} {
	if s.ElemCount == 0 {
		return nil
	}
	return s.Elements[s.ElemCount - 1]
}

func (s *intStack) Size() uint64 {
	return s.ElemCount
}

func (s *intStack) Capacity() uint64 {
	return uint64(len(s.Elements))
}
