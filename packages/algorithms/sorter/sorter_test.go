package algoritms

import (
	"github.com/stretchr/testify/require"
	"sort"
	"testing"
)

func TestQuickSort(t *testing.T) {
	arr := []int{2, 5, -4, 11, 0, 18, 22, 67, 51, 6}

	t.Run("quick sort", func(t *testing.T) {
		QuickSort(arr, 0, len(arr)-1)

		result := sort.SliceIsSorted(arr, func(i, j int) bool {
			return i < j
		})

		require.Equal(t, true, result)
	})
}
