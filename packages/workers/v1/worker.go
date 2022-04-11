package v1

import (
	"fmt"
	"sync"
)

type worker struct {
	Id       int
	TaskChan chan task

	// TODO: check
	//workerPool chan *worker
	workerPool chan chan task
	quitChan   chan bool
}

func newWorker(id int, workerPool chan chan task) *worker {
	return &worker{
		Id:       id,
		TaskChan: make(chan task),

		workerPool: workerPool,
		quitChan:   make(chan bool),
	}
}

func (w *worker) Start(wg *sync.WaitGroup) {
	fmt.Printf("Starting worker %d\n", w.Id)
	go func() {
		for {
			// Add ourselves into the worker queue.
			w.workerPool <- w.TaskChan
			select {
			case t := <-w.TaskChan:
				w.process(t)
				wg.Done()
			case <-w.quitChan:
				return
			}
		}
	}()
}

func (w *worker) Stop() {
	fmt.Printf("Stopping worker%d\n", w.Id)
	go func() {
		w.quitChan <- true
		//wg.Done()
	}()
}

func (w *worker) process(t task) {
	fmt.Printf("Invoke task by worker %d\n", w.Id)
	r := invokeTask(t)
	v := getValueFromReflectValue(r)
	t.ResultChan <- v
	fmt.Println(fmt.Sprintf("Worker %d has result %d", w.Id, v))
}
