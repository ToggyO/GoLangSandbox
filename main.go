package main

import (
	"fmt"
	"hello/packages/mapTest"
	"hello/packages/pointers"
	"hello/packages/utils"
)

func main() {
	var util = new(utils.Utils)

	util.Add(2, 4, 5, 6)
	fmt.Println(util.SumResult)

	util.Multiply(1, 2, 5, 5)
	fmt.Println(util.MultiplyResult)

	mapTest.MapTest()

	pointers.StartPointers()

	//api.StartApi()
}