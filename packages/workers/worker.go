package workers

import (
    "fmt"
    "sync"
)

type worker struct {
    Id       int
    TaskChan chan *task

    workerPool chan *worker
    quitChan   chan bool
}

func newWorker(id int, workerPool chan *worker) *worker {
    return &worker{
        Id:       id,
        TaskChan: make(chan *task, 1),

        workerPool: workerPool,
        quitChan:   make(chan bool),
    }
}

func (w *worker) Start() {
    fmt.Printf("Starting worker %d\n", w.Id)
    go func() {
        for {
            select {
            case t := <-w.TaskChan:
                r := invokeTask(t)
                t.ResultChan <- getValueFromReflectValue(r)
                w.workerPool <- w
            case <-w.quitChan:
                return
            }
        }
    }()
}

func (w *worker) Stop(wg *sync.WaitGroup) {
    fmt.Printf("Stopping worker%d\n", w.Id)
    go func() {
        w.quitChan <- true
        wg.Done()
    }()
}
