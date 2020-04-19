// Package wpool implements the Worker Pool pattern
package wpool

import (
	"fmt"
	"runtime"
	"sync"
)

// WorkerPool implementation to manage workers and jobs
type WorkerPool struct {
	maxWorkers int
	workers    []*Worker
	pool       chan chan func()
	jobs       chan func()
	quit       chan bool
	wgPool     sync.WaitGroup
	wgWorkers  *sync.WaitGroup
}

// NewWorkerPool spawn a worker pool
// if maxWorkers is 0 or below then use numcpu
func NewWorkerPool(maxWorkers int) *WorkerPool {
	if maxWorkers <= 0 {
		maxWorkers = runtime.NumCPU()
	}

	wgWorkers := sync.WaitGroup{}
	pool := make(chan chan func(), maxWorkers)
	workers := make([]*Worker, maxWorkers, maxWorkers)

	// spawn workers
	for i := 0; i < maxWorkers; i++ {
		workers[i] = NewWorker(pool, &wgWorkers)
	}
	return &WorkerPool{
		maxWorkers: maxWorkers,
		workers:    workers,
		pool:       pool,
		jobs:       make(chan func()),
		quit:       make(chan bool),
		wgPool:     sync.WaitGroup{},
		wgWorkers:  &wgWorkers,
	}
}

// Start the WorkerPool
func (wp *WorkerPool) Start() {

	// start our lovely workers
	for _, worker := range wp.workers {
		worker.Start()
	}

	go wp.dispatch()
}

// Stop the WorkerPool
func (wp *WorkerPool) Stop() {
	wp.quit <- true
	wp.wgPool.Wait()
}

// Submit a job to the WorkerPool
func (wp *WorkerPool) Submit(job func()) {
	if job != nil {
		fmt.Println("Submitting a job to the WorkerPool")
		wp.jobs <- job
	}
}

// dispatch manage the pool
func (wp *WorkerPool) dispatch() {
	wp.wgPool.Add(1)

	for {
		select {
		case job := <-wp.jobs:
			availableWorker := <-wp.pool
			availableWorker <- job
		case <-wp.quit:
			for _, worker := range wp.workers {
				worker.Stop()
			}
			wp.wgWorkers.Wait()
			wp.wgPool.Done()
			fmt.Println("WorkerPool is stopped")
			return
		}
	}
}
