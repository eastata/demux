package main

import (
	"sync"
	"testing"
)

func TestWorker(t *testing.T) {
	wg := sync.WaitGroup{}
	wg.Add(1)
	err := worker(&wg, []int{1, 2})
	if err != nil {
		t.Errorf("%v", err)
	}
	wg.Wait()
}

func BenchmarkWorker(b *testing.B) {
	for i := 0; i < b.N; i++ {
		wg := sync.WaitGroup{}
		wg.Add(1)
		worker(&wg, []int{1, 2})
		wg.Wait()
	}
}
