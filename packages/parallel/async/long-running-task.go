package async

import (
	"fmt"
	"time"
)

func RunLongRunningTask() {
	aCh, bCh, cCh := longRunningTask(), longRunningTask(), longRunningTask()
	a, b, c := <-aCh, <-bCh, <-cCh

	fmt.Println(a, b, c)
}

func longRunningTask() <-chan int32 {
	r := make(chan int32)

	go func() {
		//defer close(r)

		// Simulate a workload.
		time.Sleep(time.Second * 3)
		//r <- rand.Int31n(100)
	}()

	return r
}
