package async_academy

import "fmt"

const count = 10

func RunAsyncWrite() {
	bi := make(Bi)
	done := make(Done)

	go generatorForWriter(done, bi, count)
	go generatorForWriter(done, bi, count)
	go generatorForWriter(done, bi, count)

	go readerForWriter(bi)

	checkOnComplete(done)
	close(bi)
	fmt.Println("Done")
}

func generatorForWriter(d Done, c In, count int) {
	defer func() {
		d <- struct{}{}
	}()
	for i := 1; i <= count; i++ {
		c <- i
	}
}

func readerForWriter(c Out) {
	for v := range c {
		fmt.Println(v)
	}
}
