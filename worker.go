// Package wpool implements the Worker Pool pattern
package wpool

import (
	"log"
	"sync"
)

// Worker implementation to manage worker
type Worker struct {
	id   int
	pool chan chan Job
	jobs chan Job
	quit chan bool
	done *sync.WaitGroup
}

// NewWorker spawn a Worker
func NewWorker(id int, pool chan chan Job, done *sync.WaitGroup) *Worker {
	return &Worker{
		id:   id,
		pool: pool,
		jobs: make(chan Job),
		quit: make(chan bool),
		done: done,
	}
}

// Start job processing for worker
func (w *Worker) Start() {
	log.Printf("Worker[%v]: started", w.id)
	go func() {
		w.done.Add(1)
		for {
			w.pool <- w.jobs

			select {
			case job := <-w.jobs:
				log.Printf("Worker[%v]: is processing job [%s]", w.id, job.id)
				job.task()
				log.Printf("Worker[%v]: job [%s] processing finished", w.id, job.id)
			case <-w.quit:
				w.done.Done()
				log.Printf("Worker[%v]: stopped", w.id)
				return
			}
		}
	}()
}

// Stop job processing for worker
func (w *Worker) Stop() {
	w.quit <- true
}
