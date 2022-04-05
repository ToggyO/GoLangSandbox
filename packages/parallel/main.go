package parallel

import "hello/packages/parallel/async"

func Run() {
	//startTime := time.Now()
	//for i := 1; i < 7; i++ {
	//	factorial(i)
	//}
	//fmt.Printf("The end: %s", time.Since(startTime))
	//fmt.Scanln()

	//PushToChannel(parallel.Data{Id: 1, Name: "Something"})

	//RunStreaming()

	//RunWaitGroup()

	//async.RunAsync()

	async.RunLongRunningTask()
}
