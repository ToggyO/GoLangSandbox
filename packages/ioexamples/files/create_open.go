package ioexamples

import (
	"fmt"
	"os"
)

func CreateFile() {
	file, err := os.Create(FILE_PATH)
	if err != nil {
		fmt.Println("Unable to create file:", err)
		os.Exit(1)
	}

	defer file.Close()
	fmt.Println(file.Name())
}

func OpenFile() {
	file, err := os.OpenFile(FILE_PATH, os.O_RDONLY, 0666)
	if err != nil {
		fmt.Println("Unable to open file:", err)
		os.Exit(1)
	}

	defer file.Close()
	fmt.Println(file.Name())
}
