package poollib

import (
	"fmt"
	"github.com/gammazero/workerpool"
	"math/rand"
	"time"
)

func FillPool(pool *workerpool.WorkerPool) {
	for {
		pool.Submit(func() {
			fmt.Println("KEK")
			time.Sleep(time.Millisecond * time.Duration(1000+rand.Intn(200)))
		})
	}
}

func RunPoolLib() {
	wp := workerpool.New(2)
	//requests := []string{"alpha", "beta", "gamma", "delta", "epsilon"}

	go FillPool(wp)
	go FillPool(wp)

	//results := make(chan string, len(requests))

	//for i, r := range requests {
	//	r := r
	//	wp.Submit(func() {
	//		fmt.Println("Handling request:", r)
	//		results <- r + strconv.Itoa(i)
	//		//time.Sleep(100 * time.Millisecond)
	//	})
	//}

	wp.StopWait()

	//for r := range results {
	//	fmt.Println(r)
	//}

	//for _, r := range requests {
	//	r := r
	//	wp.Submit(func() {
	//		fmt.Println("Handling request:", r)
	//	})
	//}

	fmt.Println("RunPoolLib() end!")
}
