package workers

import (
    "sync"
)

const (
    PoolCreated = "pool-created"
    PoolRunning = "pool-running"
    PoolStopped = "pool-stopped"
)

type Pool struct {
    status string
    job    job

    taskChan   chan *task
    workerPool chan *worker
    quitChan   chan bool

    workersCount int
    wg           *sync.WaitGroup
}

func NewPool(jobFunc interface{}, workersCount int) *Pool {
    err := isFunction(jobFunc)
    if err != nil {
        panic(err)
    }

    return &Pool{
        status: PoolCreated,
        job:    newJob(jobFunc),

        taskChan: make(chan *task),

        workerPool: make(chan *worker, workersCount),
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
        w := newWorker(i, p.workerPool)
        w.Start()
        p.workerPool <- w
    }

    go func() {
        for {
            select {
            case t := <-p.taskChan:
                go func() {
                    w := <-p.workerPool
                    w.TaskChan <- t
                }()
            case <-p.quitChan:
                return
            }
        }
    }()

    p.status = PoolRunning
    // TODO: check
    //p.wg.Wait()
}

func (p *Pool) AddTask(arguments ...interface{}) <-chan interface{} {
    p.checkOnRunning()

    resultChan := make(chan interface{})
    // TODO: check
    //defer close(resultChan)

    p.taskChan <- newTask(p.job, arguments, resultChan)

    return resultChan
}

// TODO: метод не работает
func (p *Pool) Stop() {
    p.checkOnRunning()

    defer close(p.quitChan)
    defer close(p.taskChan)
    defer close(p.workerPool)

    for {
        select {
        case w := <-p.workerPool:
            p.wg.Add(1)
            w.Stop(p.wg)
        default:
            p.wg.Wait()
            p.status = PoolStopped
            p.quitChan <- true
            break
        }
    }
    //for w, ok := range p.workerPool; ok {
    //	w.Stop()
    //}
}

func (p *Pool) checkOnRunning() {
    if p.status != PoolRunning {
        panic("Call Run() method before adding tasks")
    }
}
