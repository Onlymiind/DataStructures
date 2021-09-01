package sort

import (
	"testing"
)

func Equal(lhs []int, rhs []int) bool {
	if len(lhs) != len(rhs) {
		return false
	}
	for i := range lhs {
		if lhs[i] != rhs[i] {
			return false
		}
	}
	return true
}

func SortingFuncTest(t *testing.T, sort_f func([]int) []int) {
	var empty []int = []int{}
	var sample []int = []int{10, 0, 4, 9}
	var sorted []int = []int{1, 2, 3, 4}
	var single []int = []int{1}

	empty = sort_f(empty)

	if !Equal(empty, []int{}) {
		t.Error("Wrong answer on empty sequence")
		t.Log("Result: ", empty)
	}

	sample = sort_f(sample)

	if !Equal(sample, []int{0, 4, 9, 10}) {
		t.Error("Wrong answer on common case")
		t.Log("Result: ", sample)
	}

	sorted = sort_f(sorted)

	if !Equal(sorted, []int{1, 2, 3, 4}) {
		t.Error("Wrong answer on sorted sequense")
		t.Log("Result: ", sorted)
	}

	single = sort_f(single)

	if !Equal(single, []int{1}) {
		t.Error("Wrong answer on sequence with single element")
		t.Log("Result: ", single)
	}
}

func TestInsertionSort(t *testing.T) {
	SortingFuncTest(t, InsertionSort)
}

func TestMerge(t *testing.T) {
	sample := []int{1, 2, 3, 4}
	assert_equal := func(lhs []int, rhs []int) {
		if !Equal(lhs, rhs) {
			t.Error("Merge doesn't work")
		}
	}

	result := merge([]int{1, 4}, []int{2, 3})

	assert_equal(result, sample)

	result = merge([]int{1, 2, 4}, []int{3})

	assert_equal(result, sample)

	result = merge([]int{1, 2, 3, 4}, []int{})

	assert_equal(result, sample)

}

func TestMergeSort(t *testing.T) {
	SortingFuncTest(t, MergeSort)
}
