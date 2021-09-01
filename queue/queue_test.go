package queue

import (
	"testing"
)

func TestSizeAndCapacity(t *testing.T) {
	q := NewIntQueue(10)

	if q.Size() != 0 {
		t.Error("Size of empty queue isn't 0")
	}

	if q.Capacity() != 10 {
		t.Error("Wrong capacity")
	}

	if !q.Empty() {
		t.Error("Empty() doesn't work")
	}

	q.Enqueue(1)

	if q.Size() != 1 {
		t.Error("Size or Enqueue is wrong")
	}

	if q.Empty() {
		t.Error("Empty or Enqueue is wrong")
	}
}

func TestEnqueue(t *testing.T) {
	q := NewIntQueue(10)

	for i := 0; i < 10; i++ {
		err := q.Enqueue(i)
		if err != nil {
			t.Error("Failed to Enqueue")
			break
		}
		if q.Size() != uint64(i+1) {
			t.Error("Enqueue doesn't change size")
		}
	}

	if q.Size() != q.Capacity() {
		t.Error("Queue isn't full")
	}

	err := q.Enqueue(10)

	if err == nil {
		t.Error("Pushing elements to full queue doesn't produce an error")
	}
}

func TestDequeue(t *testing.T) {
	q := NewIntQueue(10)

	err := q.Dequeue()
	if err == nil {
		t.Error("Popping element from empty queue doesn't produce an error")
	}

	q.Enqueue(1)
	err = q.Dequeue()
	if err != nil || q.Size() != 0 {
		t.Error("Dequeue doesn't work")
	}
}

func TestClear(t *testing.T) {
	q := NewIntQueue(10)

	for i := 0; i < 10; i++ {
		q.Enqueue(i)
	}
	q.Clear()

	if q.Size() != 0 {
		t.Error("Clear doesn't work")
	}
}

func TestPeek(t *testing.T) {
	q := NewIntQueue(10)

	q.Enqueue(1)
	if q.Peek() != 1 {
		t.Error("Peek doesn't work")
	}

	q.Clear()

	for i := 0; i < 10; i++ {
		q.Enqueue(i)
	}

	q.Dequeue()
	q.Enqueue(10)

	if q.Peek() != 1 {
		t.Logf("Expected 1, got %d", q.Peek())
		t.Error("Peek doesn't work when queue's tail is offset")
	}
}
