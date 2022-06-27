package algorithms

func QuickSortOptimized(source []int) {
	srt(source, 0, len(source)-1)
}

func srt(source []int, left, right int) {
	if right-left < 1 { // размер массива 1 или меньше
		return
	}

	//arr := []int{3, 5, 2, 1, 4, 8}
	// 3, 5, >2, 1, 4, 8
	// 2, 5, >3, 1, 4, 8
	// 2, 3, >5, 1, 4, 8
	//
	// 2, 3, >5, 1, 4, 8
	// 2, 3, >4, 1, 5, 8
	// 2, 3, >1, 4, 5, 8
	//
	// 2, 3, 1 l=0, r=2
	// 2, >3, 1
	// 2, >3, 1
	//
	// 2, >3, 1
	// 2, >1, 3
	//
	// >2, 1
	// >1, 2
	//
	// 3
	// [1, 2, 3]
	//
	// 4, 5, 8  l=3, r=5
	// 4, >5, 8
	// 4, >5, 8
	//
	// 4, >5, 8
	// 4, >5, 8
	//
	// 4, 5  l=3, r=4
	// >4, 5
	// >4, 5
	//
	// 8 l=5, r=5
	//
	// [4, 5, 8]
	// [1, 2, 3, 4, 5, 8]

	mid := left + (right-left)/2
	l := left
	r := right

	for l < mid {
		if source[l] > source[mid] {
			source[l], source[mid] = source[mid], source[l]
		}
		l++
	}

	for r > mid {
		if source[r] < source[mid] {
			source[r], source[mid] = source[mid], source[r]
		}
		r--
	}

	srt(source, left, mid)
	srt(source, mid+1, right)
}
