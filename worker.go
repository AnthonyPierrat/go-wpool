// Package wpool implements the Worker Pool pattern
package wpool

import (
	"fmt"
	"sync"
)

// Worker implementation to manage worker
type Worker struct {
	Pool chan chan func()
	Jobs chan func()
	quit chan bool
	done sync.WaitGroup
}

// NewWorker spawn a Worker
func NewWorker(pool chan chan func(), done sync.WaitGroup) *Worker {
	return &Worker{
		Pool: pool,
		Jobs: make(chan func()),
		quit: make(chan bool),
		done: done,
	}
}

// Start job processing for worker
func (w *Worker) Start() {
	go func() {
		w.done.Add(1)
		for {
			w.Pool <- w.Jobs

			select {
			case job := <-w.Jobs:
				fmt.Println("start processing job")
				job()
				fmt.Println("job processing finished")
			case <-w.quit:
				w.done.Done()
				return
			}
		}
	}()
}

// Stop job processing for worker
func (w *Worker) Stop() {
	w.quit <- true
}
