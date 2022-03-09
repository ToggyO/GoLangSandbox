package ioexamples

import (
	"fmt"
	"io"
	"os"
)

func ReadFile() {
	file, err := os.OpenFile(FILE_PATH, os.O_RDONLY, 0666)
	if err != nil {
		fmt.Println("Unable to create file:", err)
		os.Exit(1)
	}

	data := make([]byte, 1024)

	for {
		n, err := file.Read(data)
		if err == io.EOF {
			break
		}
		fmt.Println(string(data[:n]))
	}
}

func WriteFile() {
	file, err := os.OpenFile(FILE_PATH, os.O_APPEND, 0666)
	if err != nil {
		fmt.Println("Unable to open file:", err)
		os.Exit(1)
	}

	defer file.Close()
	file.WriteString("package ioexamples\n\nimport \"fmt\"\n\nfunc Console() {\n    fmt.Println(\"Generated .go file\")\n}")
}
