package stack

import (
	"testing"
)

func TestPush(t *testing.T) {
	s := NewIntStack(10)

	s.Push(1)

	if s.Size() != 1 {
		t.Error("Push() doesn't work")
	}
}

func TestPop(t *testing.T) {
	s := NewIntStack(10)

	s.Push(1)
	s.Pop()

	if s.Size() != 0 {
		t.Error("Pop() doesn't work")
	}
}

func TestTop(t *testing.T) {
	s := NewIntStack(10)

	s.Push(1)

	if s.Top().(int) != 1 {
		t.Error("Top() doesn't work")
	}
}

func TestCapacity(t *testing.T) {
	s := NewIntStack(10)

	if s.Capacity() != 10 {
		t.Error("Wrong capacity")
	}
}