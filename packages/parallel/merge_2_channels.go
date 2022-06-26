package parallel

import "sync"

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

		for i := 0; i < n; i++ {
			var x1 int
			var x2 int

			select {
			case x1 = <-in1:
			default:
				x1 = 0
			}

			select {
			case x2 = <-in1:
			default:
				x2 = 0
			}

			wg.Add(1)
			go func() {
				defer wg.Done()
				out <- f(x1) + f(x2)
			}()
		}

		go func() {
			wg.Wait()
			close(out)
		}()
	}(f, in1, in2, out, n)
}
