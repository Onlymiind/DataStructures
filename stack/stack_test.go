package stack

import (
	"testing"
)

func TestPush(t *testing.T) {
	s := NewIntStack(10)

	s.Push(1)
}