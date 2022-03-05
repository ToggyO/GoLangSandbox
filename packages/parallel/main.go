package parallel

import (
	parallel "hello/packages/parallel/absractions"
)

func Run() {
	//startTime := time.Now()
	//for i := 1; i < 7; i++ {
	//	factorial(i)
	//}
	//fmt.Printf("The end: %s", time.Since(startTime))
	//fmt.Scanln()

	PushToChannel(parallel.Data{Id: 1, Name: "Something"})
}
