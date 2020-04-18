package wpool

import (
	"sync"
	"testing"
)

func TestShouldCreateWorker(t *testing.T) {
	pool := make(chan chan func())
	done := sync.WaitGroup{}
	worker := NewWorker(pool, done)

	if worker == nil {
		t.Fatalf("Could not create worker")
	}
}
