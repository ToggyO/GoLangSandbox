package workers

import (
	"fmt"
	"time"
)

func add(a, b int) int {
	r := a + b
	time.Sleep(time.Duration(r*100) * time.Millisecond)
	return r
}

func calc(pool *Pool, calcOrder string) {
	for i, el := range [][]int{{1, 5}, {2, 8}, {6, 1}, {3, 8}, {5, 7}, {9, 11}} {
		result := <-pool.AddTask(el[0], el[1])
		fmt.Printf("%s calculation, iteration #%v:", calcOrder, i)
		fmt.Println(result)
	}
}

func RunPoolTest() {
	pool := NewPool(add, 5)
	pool.Run()

	time.Sleep(100 * time.Millisecond)

	//go calc(pool, "First")
	//go calc(pool, "Second")

	//time.Sleep(10000 * time.Millisecond)

	for i, el := range [][]int{{1, 5}, {2, 8}, {6, 1}, {3, 8}, {5, 7}, {9, 11}} {
		result := <-pool.AddTask(el[0], el[1])
		fmt.Printf("First calculation, iteration #%v:", i)
		fmt.Println(result)
	}

	//time.Sleep(1000 * time.Millisecond)
	//
	//for i, el := range [][]int{{1, 5}, {2, 8}, {6, 1}, {3, 8}, {5, 7}, {9, 11}} {
	//	result := pool.AddTask(el[0], el[1])
	//	fmt.Printf("Second calculation, iteration #%v:", i)
	//	fmt.Println(result)
	//}

	time.Sleep(1000 * time.Millisecond)

	pool.Stop()

	fmt.Println("End onf RunPoolTest")
}
