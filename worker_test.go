package wpool

import (
	"sync"
	"testing"
)

func TestShouldCreateWorker(t *testing.T) {
	pool := make(chan chan Job)
	done := sync.WaitGroup{}
	worker := NewWorker(1, pool, &done)

	if worker == nil {
		t.Fatalf("Could not create worker")
	}
}
