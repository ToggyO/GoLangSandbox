package v2

import (
	"context"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

type Pool struct {
	maxWorkers int

	workerQueue chan WorkerTask
	taskQueue   chan WorkerTask

	// TODO: choose between copy and pointer
	wg sync.WaitGroup
}

func NewPool(maxWorkers int) *Pool {
	p := &Pool{
		maxWorkers: maxWorkers,

		workerQueue: make(chan WorkerTask),
		taskQueue:   make(chan WorkerTask, 1),

		wg: sync.WaitGroup{},
	}

	go p.run()

	return p
}

func (p *Pool) Exec(task WorkerTask) {
	if task != nil {
		p.taskQueue <- task
	}
}

func (p *Pool) run() {
	ctx, cancel := context.WithCancel(context.Background())
	p.wg.Add(p.maxWorkers)

	defer p.wg.Done()

	// TODO: check

	//
	for i := 0; i < p.maxWorkers; i++ {
		go worker(p.workerQueue, &p.wg, i+1)
	}
	//

	for {
		select {
		case task := <-p.taskQueue:
			//TODO: check
			select {
			case <-ctx.Done():
				return
			default:
			}
			//
			p.workerQueue <- task
		case <-ctx.Done():
			close(p.workerQueue)
			return
		}
	}
	defer p.wait(cancel)
}

func (p *Pool) wait(cancel context.CancelFunc) {
	quitChan := make(chan os.Signal)
	signal.Notify(quitChan, syscall.SIGINT, syscall.SIGTERM)

	<-quitChan

	cancel()
	p.wg.Wait()
}
