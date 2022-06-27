package algorithms

func QuickSortWithBounds(source []int, low int, high int) {
	if len(source) == 0 {
		return
	}

	if low >= high {
		return
	}

	pivot := low + (high-low)/2
	left := low
	right := high

	for left <= right {
		for source[left] < source[pivot] {
			left++
		}

		for source[right] > source[pivot] {
			right--
		}

		if left <= right {
			source[left], source[right] = source[right], source[left]
			left++
			right--
		}
	}

	if low < right {
		QuickSortWithBounds(source, low, right)
	}

	if left < high {
		QuickSortWithBounds(source, left, high)
	}
}
