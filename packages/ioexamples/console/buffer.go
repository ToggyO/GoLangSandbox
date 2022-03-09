package console

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

const filePath = "./packages/ioexamples/console/buffer.dat"

func RunBuffer() {
	var utils bufferUtils

	utils.write()
	utils.read()
}

type bufferUtils struct{}

func (bf bufferUtils) write() {
	rows := []string{
		"Hello Go!",
		"Welcome to Golang",
	}

	file, err := os.Create(filePath)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	for _, row := range rows {
		writer.WriteString(row)
		writer.WriteString("\n")
	}
	writer.Flush()
}

func (bf bufferUtils) read() {
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				break
			}
			fmt.Println(err)
		}

		fmt.Println(line)
	}
}
