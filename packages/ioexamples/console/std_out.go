package console

import (
	"fmt"
	ioexamples "hello/packages/ioexamples/files"
	"io"
	"os"
)

func RunStdOutput() {
	file, err := os.Open(ioexamples.FILE_PATH)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer file.Close()

	io.Copy(os.Stdout, file)
}
