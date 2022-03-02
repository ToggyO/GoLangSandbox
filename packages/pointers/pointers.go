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
}
