package internal

import (
	"fmt"
	"log"
	"runtime/debug"
)

// WorkerPool is a concrete implementation of domain.WorkerPool for executing work against a pool of go routines
type WorkerPool struct {
	workQueue  chan func()
	numWorkers int
}

// NewWorkerPool creates a new worker pool with the specified number of go routine 'workers' and queue size
func NewWorkerPool(numWorkers, queueSize int) *WorkerPool {
	pool := &WorkerPool{
		workQueue:  make(chan func(), queueSize),
		numWorkers: numWorkers,
	}
	return pool
}

// Do adds an item of work to the work queue
func (p *WorkerPool) Do(work func()) error {
	select {
	case p.workQueue <- work:
	default:
		log.Printf("queue full")
		return fmt.Errorf("rate limited")
	}
	return nil
}

// Start creates workers to start reading from the queue
func (p *WorkerPool) Start() {
	log.Printf("Starting worker pool")
	for i := 0; i < p.numWorkers; i++ {
		go p.worker()
	}
}

// Stop closes the queue and lets the workers finish what's on there
func (p *WorkerPool) Stop() {
	log.Printf("Stopping worker pool")
	close(p.workQueue)
}

func (p *WorkerPool) worker() {
	for work := range p.workQueue {
		p.doWork(work)
	}
}

func (p *WorkerPool) doWork(work func()) {
	defer func() {
		if err := recover(); err != nil {
			log.Printf("worker recovered from panicking go routine %v %s", err, debug.Stack())
		}
	}()
	work()
}
