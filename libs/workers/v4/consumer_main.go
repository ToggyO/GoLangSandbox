package v4

import (
	"context"
	"fmt"
	"hello/libs/workers/common"
	"hello/packages/data_structures/queue"
	"sync"
	"time"
)

type consumer struct {
	workersCount int
	maxWorkers   int
	idleTimeout  time.Duration
	timeout      *time.Timer

	workerQueue chan common.WorkerTask
	taskQueue   chan common.WorkerTask

	wg           sync.WaitGroup
	ctx          context.Context
	cancel       context.CancelFunc
	waitingQueue *queue.Queue[common.WorkerTask]
	quitChan     chan bool
}

func newConsumer(maxWorkers int, ctx context.Context, cancel context.CancelFunc) consumer {
	idleTimeout := 2 * time.Second
	return consumer{
		maxWorkers:  maxWorkers,
		idleTimeout: idleTimeout,
		timeout:     time.NewTimer(idleTimeout),

		workerQueue: make(chan common.WorkerTask),
		taskQueue:   make(chan common.WorkerTask),

		ctx:          ctx,
		cancel:       cancel,
		waitingQueue: queue.NewQueue[common.WorkerTask](),
		quitChan:     make(chan bool),
	}
}

func (c *consumer) Start() {
	var idle bool

Loop:
	for {
		fmt.Println("LOOP")

		l := c.waitingQueue.Len()
		if l != 0 {
			if !c.processWaitingQueue() {
				break Loop
			}
			continue
		}

		select {
		case task, ok := <-c.taskQueue:
			if !ok {
				break Loop
			}

			select {
			case c.workerQueue <- task:
				fmt.Println("HMMMM....")
			default:
				if c.workersCount < c.maxWorkers {
					c.wg.Add(1)
					c.runWorker(task, c.workersCount+1, &c.wg)
					c.workersCount++
				} else {
					fmt.Println("Enqueue in loop")
					c.waitingQueue.Enqueue(task)
				}
			}
			idle = false
		case <-c.timeout.C:
			if idle && c.workersCount > 0 {
				if c.killIdleWorker() {
					c.workersCount--
				}
			}
			idle = true
			c.timeout.Reset(c.idleTimeout)
		case <-c.quitChan:
			break Loop
		}
	}

	c.handleStop()
}

func (c *consumer) Stop() {
	c.quitChan <- true
}
