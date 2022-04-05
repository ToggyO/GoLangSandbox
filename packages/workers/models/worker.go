package models

import (
    "fmt"
)

type Worker struct {
    Id       int
    TaskChan chan *Task

    quitChan chan bool
}

func NewWorker(id int) *Worker {
    return &Worker{
        Id:       id,
        TaskChan: make(chan *Task, 1),

        quitChan: make(chan bool),
    }
}

func (w *Worker) Start() {
    fmt.Printf("Starting worker %d\n", w.Id)
    for {
        select {
        case task := <-w.TaskChan:
            // TODO: args
            task.Execute()
        case <-w.quitChan:
            return
        }
    }
}

func (w *Worker) Stop() {
    fmt.Printf("Stopping worker%d\n", w.Id)
    go func() {
        w.quitChan <- true
    }()
}
