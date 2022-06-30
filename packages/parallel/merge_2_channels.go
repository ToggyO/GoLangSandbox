package parallel

import (
	"fmt"
	"sync"
)

func RunMerge2Channels() {
	generator := func(start, end int) <-chan int {
		result := make(chan int)

		go func() {
			for i := start; i <= end; i++ {
				result <- i
			}
		}()

		return result
	}

	multiply := func(x int) int {
		return x * x
	}

	in1 := generator(1, 5)
	in2 := generator(6, 8)
	out := make(chan int)

	Merge2Channels(multiply, in1, in2, out, 5)

	for v := range out {
		fmt.Println(v)
	}
	fmt.Println("Done")
}

func Merge2Channels(
	f func(int) int,
	in1 <-chan int,
	in2 <-chan int,
	out chan<- int,
	n int) {
	go func(f func(int) int,
		in1 <-chan int,
		in2 <-chan int,
		out chan<- int,
		n int) {
		wg := sync.WaitGroup{}
		wg.Add(n)

		for i := 0; i < n; i++ {
			// TODO: check
			x1 := <-in1
			x2 := <-in2

			go func(a, b int, r chan<- int) {
				defer wg.Done()
				r <- f(a) + f(b)
			}(x1, x2, out)
		}

		go func() {
			wg.Wait()
			close(out)
		}()
	}(f, in1, in2, out, n)
}
