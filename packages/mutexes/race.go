package mutexes

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func Run() {

	var counter int
	var mu sync.Mutex

	fmt.Println("Initial value: ", counter)

	// deploy 5 goroutines
	for i := 0; i < 5; i++ {

		wg.Add(1)

		go func() {
			defer wg.Done()
			//increment the counter 100 times
			for j := 0; j < 100; j++ {
				mu.Lock()

				temp := counter
				time.Sleep(time.Microsecond * 1)
				temp += 1
				counter = temp

				mu.Unlock()
			}

		}()

	}

	wg.Wait()
	fmt.Println("Final value: ", counter)
}
