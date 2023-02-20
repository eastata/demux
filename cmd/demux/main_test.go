package main

import (
	"github.com/google/uuid"
	"testing"
)

func TestJobsGenerator(t *testing.T) {
	jobs := jobsGenerator()
	if jobs[0].id == jobs[2].id {
		t.Errorf("Collision for generated jobs UUID")
	}
}

func BenchmarkJonsGenerator(b *testing.B) {
	for i := 0; i < b.N; i++ {
		jobsGenerator()
	}
}

func TestShceduler(t *testing.T) {
	jobs := jobsGenerator()
	scheduler(jobs)
}

func BenchmarkScheduler(b *testing.B) {
	jobs := jobsGenerator()
	for i := 0; i < b.N; i++ {
		scheduler(jobs)
	}
}

func TestWorker(t *testing.T) {
	jb := job{id: uuid.New(), data: []int{1, 2}}

	ch := make(chan jobResponse, 1)

	err := worker(ch, jb)
	if err != nil {
		t.Errorf("%v", err)
	}
	close(ch)
}

func BenchmarkWorker(b *testing.B) {
	jb := job{id: uuid.New(), data: []int{1, 2}}
	ch := make(chan jobResponse)

	go func() {
		for i := 0; i < b.N; i++ {
			<-ch
		}
	}()

	for i := 0; i < b.N; i++ {
		worker(ch, jb)

	}
}
