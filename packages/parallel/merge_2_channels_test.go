package parallel

import (
	"fmt"
	"github.com/stretchr/testify/require"
	"testing"
)

func generator(start, end int) <-chan int {
	result := make(chan int)

	go func() {
		for i := start; i < end; i++ {
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

		Merge2Channels(multiply, in1, in2, out, 5)

		for v := range out {
			fmt.Println(v)
		}

		require.NotEmpty(t, channel)
	})
}
