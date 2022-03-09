package console

import (
	"fmt"
	"os"
)

type Person struct {
	Name   string
	Age    int
	Weight float32
}

func RunScan() {
	filePath := "./packages/ioexamples/files/scan_test.txt"
	writeData(filePath)
	readData(filePath)
}

func writeData(filePath string) {
	name := "Tom"
	age := 24

	file, err := os.Create(filePath)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer file.Close()

	fmt.Fprintln(file, name)
	fmt.Fprintln(file, age)
}

func readData(filePath string) {
	var name string
	var age int

	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer file.Close()

	fmt.Fscanln(file, &name)
	fmt.Fscanln(file, &age)
	fmt.Println(name, age)
}
