package homework

import (
	"fmt"
	"sync"
)

type Task func() error

func Run(tasks []Task, maxWorkersCount, maxErrorsCount int) error {
	workersCount := calcWorkersCount(len(tasks), maxWorkersCount)

	wg := &sync.WaitGroup{}
	stopChan := make(chan struct{})
	errorChan := make(chan struct{}, maxErrorsCount)
	workerQueue := make(chan Task, workersCount)

	var isErrorLimitOverflow bool

	for i := 1; i <= workersCount; i++ {
		wg.Add(1)
		go worker(i, workerQueue, errorChan, stopChan, wg)
	}

	go handleError(maxErrorsCount, errorChan, stopChan, &isErrorLimitOverflow)
	go pushTasks(tasks, workerQueue)

	//<-stopChan

	wg.Wait()
	return nil
}

func worker(id int,
	workerQueue <-chan Task,
	errorChan chan struct{},
	stopChan <-chan struct{},
	wg *sync.WaitGroup,
) {
	defer wg.Done()
	fmt.Printf("Worker with id=%d started\n", id)

	for t := range workerQueue {
		select {
		case _, ok := <-stopChan:
			if !ok {
				break
			}
		default:
			if err := t(); err != nil {
				errorChan <- struct{}{}
			}
		}
	}

	fmt.Printf("Worker with id=%d interrupted\n", id)
}

func calcWorkersCount(tasksCount, maxWorkersCount int) int {
	workersCount := maxWorkersCount
	if tasksCount < maxWorkersCount {
		workersCount = tasksCount
	}

	return workersCount
}

func pushTasks(tasks []Task, workerQueue chan<- Task) {
	for _, t := range tasks {
		workerQueue <- t
	}
	close(workerQueue)
}

func handleError(maxErrs int, errorChan chan struct{}, stopChan chan struct{}, isErrorLimitOverflow *bool) {
	var errorsCount int32
	for {
		_, ok := <-errorChan
		if ok {
			errorsCount++
		}

		if errorsCount >= int32(maxErrs) {
			*isErrorLimitOverflow = true
			close(stopChan)
			return
		}
	}
}
