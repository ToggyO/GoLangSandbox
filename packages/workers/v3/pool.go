package v3

import (
	"context"
	"fmt"
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

	stopped     bool
	stoppedChan chan struct{}
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

		stoppedChan: make(chan struct{}),
	}

	go p.run()
	go p.handleShutdown()

	return p
}

func (p *Pool) run() {
	// TODO: check possibility to add delta during worker processing his job
	p.wg.Add(p.maxWorkers)

	for i := 0; i < p.maxWorkers; i++ {
		go p.consumer.runWorker(i+1, p.wg)
	}

	p.consumer.start(p.ctx)
}

func (p *Pool) Exec(task WorkerTask) {
	if !p.stopped && task != nil {
		p.consumer.taskQueue <- task
	}
}

func (p *Pool) Watch() {
	<-p.stoppedChan
}

func (p *Pool) handleShutdown() {
	defer close(p.consumer.taskQueue)
	defer close(p.stoppedChan)

	quitChan := make(chan os.Signal)
	signal.Notify(quitChan, syscall.SIGINT, syscall.SIGTERM)

	<-quitChan

	p.cancel()
	p.wg.Wait()

	// TODO: delete
	fmt.Println("After wait")
}
