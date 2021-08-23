package list

import (
	"fmt"
	"testing"
)

func TestEmpty(t *testing.T) {
	l := NewList()
	if !l.Empty() {
		t.Error("New list isn't empty")
	}
}

func TestPushBack(t *testing.T) {
	l := NewList()
	l.PushBack(1)
	if l.Size() != 1 {
		t.Error("PushBack doesn't work")
	}
}

func TestPushFront(t *testing.T) {
	l := NewList()
	l.PushFront(1)
	if l.Size() != 1 {
		t.Error("PushFront doesn't work")
	}
}

func TestInsert(t *testing.T) {
	l:= NewList()
	l.Insert(l.End(), 1)

	if l.Size() != 1 {
		t.Error("Insert doesn't work")
	}
}

func TestErase(t *testing.T) {
	l := NewList()
	l.PushBack(1)
	l.Erase(l.Begin())

	if l.Size() != 0 {
		t.Error("Erase doesn't work")
	}
}

func TestClear(t *testing.T) {
	l := NewList()
	l.PushBack(1)
	l.Clear()
	if l.Size() != 0 {
		t.Error("Clear doesn't work")
	}
}

func TestIteration(t *testing.T) {
	l := NewList()
	if !l.Begin().Equal(l.End()) {
		t.Error("Wrong iterator comparison")
	}
	var sample [8]int = [8]int{1, 2, 3, 4, 5, 6, 7, 8}
	for _, val := range sample {
		l.PushBack(val)
	}

	for i, it := 0, l.Begin(); i < len(sample); i++ {
		if it.Get() != sample[i] {
			t.Error("Iteration doesn't work")
		}
		it.Inc()
	}

	l.Clear()

	//Do the same for PushFront
	for _, val := range sample {
		l.PushFront(val)
	}

	for i, it := 0, l.Begin(); i < len(sample); i++ {
		if it.Get() != sample[len(sample) - i - 1] {
			t.Error("Iteration doesn't work")
		}
		it.Inc()
	}

	l.Clear()

	//And for Insert
	for _, val := range sample {
		l.Insert(l.End(), val)
	}

	for i, it := 0, l.Begin(); i < len(sample); i++ {
		if it.Get() != sample[i] {
			fmt.Println(it.Get(), sample[i])
			t.Error("Iteration doesn't work")
		}
		it.Inc()
	}
}

