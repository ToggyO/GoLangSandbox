package v4

import (
    "context"
    "errors"
    "fmt"
    "hello/libs/workers/common"
    "os"
    "os/signal"
    "syscall"
)

type Pool struct {
    consumer consumer

    ctx    context.Context
    cancel context.CancelFunc

    awaitChan chan struct{}
    stopChan  chan struct{}
}

func NewPool(maxWorkers int) *Pool {
    ctx, cancel := context.WithCancel(context.Background())

    p := &Pool{
        consumer: newConsumer(maxWorkers, ctx, cancel),

        ctx:    ctx,
        cancel: cancel,

        awaitChan: make(chan struct{}),
        stopChan:  make(chan struct{}),
    }

    go p.consumer.Process()
    go p.handleDispose()

    return p
}

func (p *Pool) Exec(task common.WorkerTask) error {
    if task != nil && common.SafeSend[common.WorkerTask](p.consumer.taskQueue, task) {
        return errors.New("cannot send to disposed pool")
    }
    return nil
}

func (p *Pool) Watch() {
    <-p.awaitChan
}

func (p *Pool) Stop() {
    defer func() {
        fmt.Println("End of pool Stop()")
    }()

    p.stop()
    p.awaitChan <- struct{}{}
}

func (p *Pool) stop() {
    p.consumer.Stop()
    <-p.ctx.Done()
    fmt.Println("End of pool stop()")
}

func (p *Pool) handleDispose() {
    defer close(p.consumer.taskQueue)
    defer close(p.awaitChan)

    syscallChan := make(chan os.Signal)
    signal.Notify(syscallChan, syscall.SIGINT, syscall.SIGTERM)

    <-syscallChan
    p.stop()
    fmt.Println("End of handleDispose()")
}
