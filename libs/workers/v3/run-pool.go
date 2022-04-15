package v3

import (
	"fmt"
	"math/rand"
	"time"
)

func fillPool(pool *Pool, results chan string) {
	for {
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

func RunPool() {
	fmt.Println("<================ RunPool() start ================>")

	pool := NewPool(5)

	results := make(chan string)

	go fillPool(pool, results)
	go fillPool(pool, results)

	go func() {
		for r := range results {
			fmt.Println(r)
		}
	}()

	pool.Watch()

	//time.Sleep(100000 * time.Millisecond)

	fmt.Println("<================ RunPool() end ================>")
}
