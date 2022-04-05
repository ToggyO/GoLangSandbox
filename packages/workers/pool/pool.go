package workers

import (
    "hello/packages/workers/models"
    "sync"
)

const (
    PoolCreated = "pool-created"
    PoolRunning = "pool-running"
    PoolStopped = "pool-stopped"
)

type Pool struct {
    status string

    taskChan   chan *models.Task
    workerPool chan *models.Worker
    quitChan   chan bool

    workersCount int
    wg           *sync.WaitGroup
}

func NewPool(workersCount int) *Pool {
    return &Pool{
        status: PoolCreated,

        taskChan: make(chan *models.Task),

        workerPool: make(chan *models.Worker, workersCount),
        quitChan:   make(chan bool),

        workersCount: workersCount,
        // TODO: check
        wg: &sync.WaitGroup{},
    }
}

func (p *Pool) Run() {
    if p.status == PoolRunning {
        return
    }

    for i := 1; i <= p.workersCount; i++ {
        worker := models.NewWorker(i)
        worker.Start()
        p.workerPool <- worker
    }

    go func() {
        for {
            select {
            case task := <-p.taskChan:
                worker := <-p.workerPool
                worker.TaskChan <- task
            }
        }
    }()

    p.status = PoolRunning
    // TODO: check
    p.wg.Wait()
}

func (p *Pool) AddTask(task *models.Task) {
    if p.status != PoolRunning {
        panic("Call Run() method before adding tasks")
    }
    p.taskChan <- task
}

// TODO: implement stop
