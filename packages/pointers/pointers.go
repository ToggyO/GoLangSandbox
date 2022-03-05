package pointers

import "fmt"

func StartPointers() {
	var k = 5
	var p = &k

	fmt.Println("Address:", p)
	fmt.Println("Value:", *p)

	fmt.Println("Transforming variable...")
	k = 10
	fmt.Println("Value:", *p)

	fmt.Println("Transforming pointer...")
	*p = 15
	fmt.Println("Value:", *p)

	doWork()
}

func doWork() {
	d := 5
	fmt.Println("d before:", d)

	changeValue(&d)
	fmt.Println("d after:", d)

	createPointer(d)
}

func changeValue(x *int) {
	*x = (*x) * (*x)
}

func createPointer(x int) *int {
	p := new(int)
	*p = x
	return p
}
