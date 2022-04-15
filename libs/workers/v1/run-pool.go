package v1

import (
	"fmt"
	"time"
)

func add(a, b int) int {
	r := a + b
	//time.Sleep(time.Duration(r*1000) * time.Millisecond)
	time.Sleep(1000 * time.Millisecond)
	return r
}

func calc(pool *Pool, calcOrder string) {
	for _, el := range [][]int{{1, 5}, {2, 8}, {6, 1}, {3, 8}, {5, 7}, {9, 11}} {
		//result := <-pool.AddTask(el[0], el[1])
		pool.AddTask(el[0], el[1])
		//fmt.Printf("%s calculation, iteration #%v:", calcOrder, i)
		//fmt.Println(result)
	}
}

func RunPoolTest() {
	pool := NewPool(add, 5)
	pool.Run()

	go calc(pool, "")

	//time.Sleep(100 * time.Millisecond)

	//go calc(pool, "First")
	//go calc(pool, "Second")

	//time.Sleep(10000 * time.Millisecond)

	//for _, el := range [][]int{{1, 5, 69}, {2, 8, 70}, {6, 1, 71}, {3, 8, 72}, {5, 7, 73}, {9, 11, 74}} {
	//	//result := <-pool.AddTask(el[0], el[1])
	//	pool.AddTask(el[0], el[1])
	//	//fmt.Println(fmt.Sprintf("First calculation, iteration #%v - %v:", i, el[2]))
	//	//fmt.Println(result)
	//}
	//
	//time.Sleep(1000 * time.Millisecond)
	//
	//for _, el := range [][]int{{1, 5}, {2, 8}, {6, 1}, {3, 8}, {5, 7}, {9, 11}} {
	//	pool.AddTask(el[0], el[1])
	//	//fmt.Printf("Second calculation, iteration #%v:", i)
	//	//fmt.Println(result)
	//}

	//time.Sleep(10000000 * time.Millisecond)

	//pool.Stop()

	fmt.Println("End onf RunPoolTest")
}
