package v4

import (
	"fmt"
)

func (c *consumer) killIdleWorker() bool {
	select {
	case c.workerQueue <- nil:
		// Sent kill signal to worker.
		return true
	default:
		// No ready workers.  All, if any, workers are busy.
		return false
	}
}

func (c *consumer) processWaitingQueue() bool {
	select {
	case task, ok := <-c.taskQueue:
		if !ok {
			return false
		}
		fmt.Println("Waiting queue before Enqueue() size: ", c.waitingQueue.Len())
		c.waitingQueue.Enqueue(task)
		fmt.Println("Waiting queue after Enqueue() size: ", c.waitingQueue.Len())
	case c.workerQueue <- c.waitingQueue.Peek():
		fmt.Println("Waiting queue size before Dequeue(): ", c.waitingQueue.Len())
		_ = c.waitingQueue.Dequeue()
		fmt.Println("Waiting queue size after Dequeue(): ", c.waitingQueue.Len())
	}

	return true
}

func (c *consumer) handleStop() {
	// Stop all remaining workers as they become ready.
	for c.workersCount > 0 {
		c.workerQueue <- nil
		c.workersCount--
	}
	fmt.Println("consumer before wg.Wait()")
	c.wg.Wait()
	fmt.Println("consumer after wg.Wait()")

	c.timeout.Stop()
	c.cancel()
	fmt.Println("consumer end of run()")
}
