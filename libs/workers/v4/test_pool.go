package v4

import (
	"fmt"
	"math/rand"
	"time"
)

func fillPool(pool *Pool, results chan string) {
	for i := 0; i <= 30; i++ {
		err := pool.Exec(func() {
			fmt.Println("KEK")
			time.Sleep(time.Millisecond * time.Duration(1000+rand.Intn(200)))
			results <- "SHPEK"
		})
		if err != nil {
			break
		}
	}
}

func RunPool(rageQuit bool) {
	fmt.Println("<================ RunPool() start ================>")

	pool := NewPool(10)

	results := make(chan string)

	go fillPool(pool, results)
	go fillPool(pool, results)

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
