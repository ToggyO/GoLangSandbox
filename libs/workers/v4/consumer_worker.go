package v4

import (
	"fmt"
	"hello/libs/workers/common"
	"math/rand"
	"sync"
	"time"
)

func (c *consumer) runWorker(task common.WorkerTask, id int, wg *sync.WaitGroup) {
	task()
	go c.worker(id, wg)
}

func (c *consumer) worker(id int, wg *sync.WaitGroup) {
	for task := range c.workerQueue {
		if task == nil {
			fmt.Printf("Worker %d interrupted\n", id)
			wg.Done()
			return
		}

		fmt.Printf("Worker %d started processing job\n", id)
		task()
		// TODO: remove
		time.Sleep(time.Millisecond * time.Duration(1000+rand.Intn(2000)))
		fmt.Printf("Worker %d finished processing job\n", id)
	}
}
