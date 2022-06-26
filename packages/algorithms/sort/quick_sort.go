package algorithms

func QuickSort(source []int) []int {
	length := len(source)
	if length <= 1 {
		return source
	}

	mid := (length - 1) / 2
	// TODO: check
	var left []int
	var right []int

	for i := 0; i < length; i++ {
		if i == mid {
			continue
		}

		if source[i] > source[mid] {
			right = append(right, source[i])
		} else {
			left = append(left, source[i])
		}
	}

	var result []int

	result = append(result, QuickSort(left)...)
	result = append(result, source[mid])
	result = append(result, QuickSort(right)...)

	return result
}
