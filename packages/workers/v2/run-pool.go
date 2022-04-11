package v2

import "fmt"

func RunPool() {
	pool := NewPool(2)

	requests := []string{"alpha", "beta", "gamma", "delta", "epsilon"}
	results := make(chan string, len(requests))

	for _, r := range requests {
		r := r
		pool.Exec(func() {
			results <- r
		})
	}

	//for r := range results {
	//	fmt.Println(r)
	//}
	fmt.Println("RunPool() End")
}
