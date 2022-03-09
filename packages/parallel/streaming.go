package parallel

import "fmt"

func RunStreaming() {
	intCh := make(chan int)

	go factorialStream(7, intCh)

	for num := range intCh {
		fmt.Println(num)
	}
}

func factorialStream(n int, ch chan int) {
	defer close(ch)
	result := 1

	for i := 1; i <= n; i++ {
		result *= i
		ch <- result
	}
}
