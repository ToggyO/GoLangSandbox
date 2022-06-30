package parallel

import (
	"fmt"
	"sync"
)

var lock sync.Mutex

func Merge2ChannelsV3(
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
		lock.Lock()

		inSlice1 := make([]*int, n)
		inSlice2 := make([]*int, n)

		go handleInput(in1, inSlice1, n, f)
		go handleInput(in2, inSlice2, n, f)

		go handleMerge(inSlice1, inSlice2, out, n)
	}(f, in1, in2, out, n)
}

func handleInput(inChan <-chan int, intSlice []*int, num int, f func(int) int) {
	for i := 0; i < num; i++ {
		x := <-inChan
		go func(i, x int) {
			r := f(x)
			intSlice[i] = &r
		}(i, x)
	}
	fmt.Println("done")
}

func handleMerge(
	inputSlice1 []*int,
	inputSlice2 []*int,
	output chan<- int,
	iterationsCount int,
) {
	defer close(output)

	var i int
	for {
		if i == iterationsCount {
			lock.Unlock()
			break
		}

		x1ptr := inputSlice1[i]
		x2ptr := inputSlice2[i]
		if x1ptr != nil && x2ptr != nil {
			output <- *x1ptr + *x2ptr
			i++
		}
	}
}
