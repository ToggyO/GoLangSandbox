package async_academy

import "fmt"

func RunAsyncRead() {
	const count = 30

	bi := make(Bi)
	done := make(Done)

	go readerForReader(done, bi, "s1")
	go readerForReader(done, bi, "s2")
	go readerForReader(done, bi, "s3")

	go generatorForReader(bi, count)

	checkOnComplete(done)
	fmt.Println("Done")
}

func generatorForReader(c In, count int) {
	defer close(c)
	for i := 1; i <= count; i++ {
		c <- i
	}
}

func readerForReader(d Done, c Out, s string) {
	defer func() {
		d <- struct{}{}
		fmt.Printf("%s done \n", s)

	}()
	for v := range c {
		fmt.Printf("%s %d\n", s, v)
	}
}
