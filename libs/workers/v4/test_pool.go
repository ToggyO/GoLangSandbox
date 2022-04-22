package v4

import (
	"fmt"
	"math/rand"
	"time"
)

func fillPool(pool *Pool, results chan string, taskIdentifier string) {
	for i := 0; i <= 15; i++ {
		err := pool.Exec(func() {
			time.Sleep(time.Millisecond * time.Duration(1000+rand.Intn(200)))
			results <- fmt.Sprintf("Set result of [%s] #%d", taskIdentifier, i)
		})
		if err != nil {
			break
		}
	}
}

func RunPool(rageQuit bool) {
	fmt.Println("<================ RunPool() start ================>")

	pool := NewPool(2)

	results := make(chan string)

	go fillPool(pool, results, "First fill pool task")
	go fillPool(pool, results, "Second fill pool task")

	go func() {
		for r := range results {
			fmt.Println(r)
		}
	}()

	if rageQuit {
		go func() {
			time.Sleep(10 * time.Second)
			pool.Stop()
		}()
	}

	pool.Watch()

	fmt.Println("<================ RunPool() end ================>")
}
