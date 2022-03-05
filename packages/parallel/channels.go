package parallel

import (
	"fmt"
	parallel "hello/packages/parallel/absractions"
)

var channel = make(chan parallel.Data)

func PushToChannel(data parallel.Data) {
	go func() {
		fmt.Println("Go routine starts")
		channel <- data
	}()

	fmt.Println(<-channel)
	fmt.Println("The end")
}
