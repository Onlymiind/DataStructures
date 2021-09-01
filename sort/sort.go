package sort

func InsertionSort(elems []int) []int {
	for i := 1; i < len(elems); i++ {
		for j := i - 1; j >= 0; j-- {
			if elems[i] < elems[j] {
				elems[i], elems[j] = elems[j], elems[i]
			} else {
				break
			}
		}
	}
	return elems
}

func merge(lhs []int, rhs []int) []int {
	result := make([]int, 0, len(rhs)+len(lhs))
	lhs_first := 0
	rhs_first := 0

	for i := 0; lhs_first < len(lhs) && rhs_first < len(rhs); i++ {
		if lhs[lhs_first] < rhs[rhs_first] {
			result = append(result, lhs[lhs_first])
			lhs_first++
		} else {
			result = append(result, rhs[rhs_first])
			rhs_first++
		}
	}

	if lhs_first < len(lhs) {
		result = append(result, lhs[lhs_first:]...)
	} else if rhs_first < len(rhs) {
		result = append(result, rhs[rhs_first:]...)
	}

	return result
}

func MergeSort(elems []int) []int {
	if len(elems) == 1 || len(elems) == 0 {
		return elems
	}

	middle := len(elems) / 2

	return merge(MergeSort(elems[:middle]), MergeSort(elems[middle:]))
}
