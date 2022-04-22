package v3

import (
	"context"
	"fmt"
	"hello/libs/workers/common"
	"math/rand"
	"sync"
	"time"
)

type consumer struct {
	workerQueue chan common.WorkerTask
	taskQueue   chan common.WorkerTask
}

func newConsumer(maxWorkers int) consumer {
	return consumer{
		workerQueue: make(chan common.WorkerTask, maxWorkers),
		taskQueue:   make(chan common.WorkerTask, 1),
	}
}

func (c *consumer) start(ctx context.Context) {
	for {
		select {
		case task := <-c.taskQueue:
			c.workerQueue <- task
		case <-ctx.Done():
			fmt.Println("Consumer received cancellation signal, closing jobsChan!")
			close(c.workerQueue)
			fmt.Println("Consumer closed jobsChan")
			return
		}
	}
}

func (c *consumer) runWorker(id int, wg *sync.WaitGroup) {
	defer wg.Done()

	for task := range c.workerQueue {
		fmt.Printf("Worker %d started processing job\n", id)
		task()
		// TODO: remove
		time.Sleep(time.Millisecond * time.Duration(1000+rand.Intn(2000)))
		fmt.Printf("Worker %d finished processing job\n", id)
	}

	fmt.Printf("Worker %d interrupted\n", id)
}
