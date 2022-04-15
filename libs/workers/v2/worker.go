package v2

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func worker(workerQueue chan WorkerTask, wg *sync.WaitGroup, id int) {
	for task := range workerQueue {
		if task == nil {
			wg.Done()
			return
		}

		fmt.Printf("Worker %d started job\n", id)
		task()
		time.Sleep(time.Millisecond * time.Duration(1000+rand.Intn(2000)))
		fmt.Printf("Worker %d finished processing job\n", id)
	}
}
