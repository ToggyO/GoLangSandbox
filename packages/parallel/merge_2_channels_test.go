package parallel

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func generator(start, end int) <-chan int {
	result := make(chan int)

	go func() {
		defer close(result)
		for i := start; i <= end; i++ {
			result <- i
		}
	}()

	return result
}

func TestMerge2Channels(t *testing.T) {
	t.Run("basic", func(t *testing.T) {
		multiply := func(x int) int {
			return x * x
		}

		in1 := generator(1, 5)
		in2 := generator(6, 10)
		out := make(chan int)

		Merge2ChannelsV3(multiply, in1, in2, out, 5)

		results := make([]int, 0, 5)
		for v := range out {
			results = append(results, v)
		}

		require.Equal(t, []int{37, 53, 73, 97, 125}, results)
	})
}
