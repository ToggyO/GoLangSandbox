package algorithms

import (
	"fmt"
	"github.com/stretchr/testify/require"
	"sort"
	"testing"
)

func TestQuickSort(t *testing.T) {
	t.Run("quick sort test", func(t *testing.T) {
		arr := []int{3, 5, 2, 1, 4, 8}

		result := QuickSort(arr)

		fmt.Println(result)
		require.True(t, sort.SliceIsSorted(result, func(i, j int) bool {
			return i < j
		}))
	})

	t.Run("quick sort optimized test", func(t *testing.T) {
		arr := []int{3, -4, 5, 2, 1, -10, 4, 8, 9, -2}

		QuickSortOptimized(arr)

		fmt.Println(arr)
		require.True(t, sort.SliceIsSorted(arr, func(i, j int) bool {
			return i < j
		}))
	})

	t.Run("quick sort with bounds", func(t *testing.T) {
		arr := []int{2, 5, -4, 11, 0, 18, 22, 67, 51, 6}

		QuickSortWithBounds(arr, 0, len(arr)-1)

		result := sort.SliceIsSorted(arr, func(i, j int) bool {
			return i < j
		})

		require.Equal(t, true, result)
	})
}
