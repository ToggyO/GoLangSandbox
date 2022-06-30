package main

import "hello/packages/parallel"

func main() {
	//generics.Runtime()
	//var util = new(utils.Utils)
	//
	//util.Add(2, 4, 5, 6)
	//fmt.Println(util.SumResult)
	//
	//util.Multiply(1, 2, 5, 5)
	//fmt.Println(util.MultiplyResult)
	//
	//mapTest.MapTest()
	//
	//pointers.StartPointers()

	//api.StartApi()

	//middlewareexample.Run()

	//mutexes.Run()

	//interfaces.Run()

	//parallel.Run()

	//ioexamples.Run()

	//network.Run()

	//operations.RunOperations()

	//v1.RunPoolTest()
	//v2.RunPool()
	//v3.RunPool()
	//poollib.RunPoolLib()
	//poollib.RunPoolTest1()
	//data_structures.RunTestDataStructures()
	//v4.RunPool(false)
	//strings_test.RunStrings()
	//mutexes.RunCond()
	parallel.RunMerge2Channels()
	//v4.RunPool(false)
}

// TODO: delete
//func main() {
//	messages := make(chan string)
//	signals := make(chan bool)
//
//	select {
//	case msg := <-messages:
//		fmt.Println("received message", msg)
//	default:
//		fmt.Println("no message received")
//	}
//
//	msg := "hi"
//	select {
//	case messages <- msg:
//		fmt.Println("sent message", msg)
//	default:
//		fmt.Println("no message sent")
//	}
//
//	select {
//	case msg := <-messages:
//		fmt.Println("received message", msg)
//	case sig := <-signals:
//		fmt.Println("received signal", sig)
//	default:
//		fmt.Println("no activity")
//	}
//}
