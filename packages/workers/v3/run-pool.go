package v3

import (
	"fmt"
	"math/rand"
	"time"
)

func FillPool(pool *Pool) {
	for {
		pool.Exec(func() {
			fmt.Println("KEK")
			time.Sleep(time.Millisecond * time.Duration(1000+rand.Intn(200)))
		})
	}
}

func RunPool() {
	pool := NewPool(5)

	go FillPool(pool)
	go FillPool(pool)

	pool.Watch()

	//requests := []string{"alpha", "beta", "gamma", "delta", "epsilon"}
	//results := make(chan string, len(requests))
	//
	//for _, r := range requests {
	//	r := r
	//	pool.Exec(func() {
	//		results <- r
	//		//fmt.Println(r)
	//	})
	//}

	//for r := range results {
	//	fmt.Println(r)
	//}

	fmt.Println("RunPool() End")
}
