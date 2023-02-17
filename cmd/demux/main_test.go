package main

import (
	"github.com/google/uuid"
	"sync"
	"testing"
)

func TestWorker(t *testing.T) {
	jb := job{id: uuid.New(), data: []int{1, 2}}

	wg := sync.WaitGroup{}
	wg.Add(1)
	err := worker(&wg, jb)
	if err != nil {
		t.Errorf("%v", err)
	}
	wg.Wait()
}

func BenchmarkWorker(b *testing.B) {
	jb := job{id: uuid.New(), data: []int{1, 2}}
	for i := 0; i < b.N; i++ {
		wg := sync.WaitGroup{}
		wg.Add(1)
		worker(&wg, jb)
		wg.Wait()
	}
}
