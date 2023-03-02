package demux

import (
	"github.com/google/uuid"
	"testing"
)

func TestJobsGenerator(t *testing.T) {
	jobs := JobsGenerator()
	if jobs[0].Id == jobs[2].Id {
		t.Errorf("Collision for generated jobs UUID")
	}
}

func BenchmarkJonsGenerator(b *testing.B) {
	for i := 0; i < b.N; i++ {
		JobsGenerator()
	}
}

func TestShceduler(t *testing.T) {
	jobs := JobsGenerator()
	Scheduler(jobs)
}

func BenchmarkScheduler(b *testing.B) {
	jobs := JobsGenerator()
	for i := 0; i < b.N; i++ {
		Scheduler(jobs)
	}
}

func TestWorker(t *testing.T) {
	jb := Job{Id: uuid.New(), data: []int64{1, 2}}

	ch := make(chan jobResponse, 1)

	err := worker(ch, jb)
	if err != nil {
		t.Errorf("%v", err)
	}
	close(ch)
}

func BenchmarkWorker(b *testing.B) {
	jb := Job{Id: uuid.New(), data: []int64{1, 2}}
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
