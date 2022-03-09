package console

import (
	"fmt"
	"os"
)

type person struct {
	name   string
	age    int32
	weight float64
}

func RunFScan() {
	filename := "./packages/ioexamples/files/fscan_test.dat"
	writeDataFscan(filename)
	readDataFscan(filename)
}

func writeDataFscan(filename string) {
	// начальные данные
	tom := person{name: "Tom", age: 24, weight: 68.5}

	file, err := os.Create(filename)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer file.Close()

	// сохраняем данные в файл
	fmt.Fprintf(file, "%s %d %.2f\n", tom.name, tom.age, tom.weight)
}

func readDataFscan(filename string) {

	// переменные для считывания данных
	var name string
	var age int
	var weight float64

	file, err := os.Open(filename)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer file.Close()

	// считывание данных из файла
	_, err = fmt.Fscanf(file, "%s %d %f\n", &name, &age, &weight)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Printf("%-8s %-8d %-8.2f\n", name, age, weight)
}
