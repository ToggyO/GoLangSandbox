package v3

import (
    "context"
    "errors"
    "os"
    "os/signal"
    "sync"
    "syscall"
)

type Pool struct {
    maxWorkers int

    consumer consumer

    ctx    context.Context
    cancel context.CancelFunc
    wg     *sync.WaitGroup

    stopped   bool
    awaitChan chan struct{}
}

func NewPool(maxWorkers int) *Pool {
    ctx, cancel := context.WithCancel(context.Background())
    c := newConsumer(maxWorkers)

    p := &Pool{
        maxWorkers: maxWorkers,

        consumer: c,

        ctx:    ctx,
        cancel: cancel,
        wg:     &sync.WaitGroup{},

        awaitChan: make(chan struct{}),
    }

    go p.run()
    go p.handleDispose()

    return p
}

func (p *Pool) Exec(task WorkerTask) error {
    if task != nil && safeSend[WorkerTask](p.consumer.taskQueue, task) {
        return errors.New("cannot send to disposed pool")
    }
    return nil
}

func (p *Pool) Watch() {
    <-p.awaitChan
}

func (p *Pool) run() {
    // TODO: check
    // TODO: check possibility to add delta during worker processing his job
    //p.wg.Add(p.maxWorkers)

    for i := 0; i < p.maxWorkers; i++ {
        // TODO: check
        p.wg.Add(1)
        go p.consumer.runWorker(i+1, p.wg)
    }

    p.consumer.start(p.ctx)
}

func (p *Pool) handleDispose() {
    defer close(p.consumer.taskQueue)
    defer close(p.awaitChan)

    quitChan := make(chan os.Signal)
    signal.Notify(quitChan, syscall.SIGINT, syscall.SIGTERM)

    <-quitChan

    p.cancel()
    p.wg.Wait()
}
