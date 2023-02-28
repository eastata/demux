package demux

import (
	"fmt"
	"github.com/google/uuid"
	"math/rand"
	"runtime"
)

type Job struct {
	id   uuid.UUID
	data []int64
}

type jobResponse struct {
	id  uuid.UUID
	out int64
}

func NewJob(data []int64) Job {
	var (
		jb Job
		id uuid.UUID
	)
	uuid.EnableRandPool()

	id = uuid.New()
	jb = Job{id, data}
	return jb

}

func JobsGenerator() []Job {
	var (
		jb Job
		id uuid.UUID
	)

	// uuid.EnableRandPool() are not thread-safe and should only be called when there is no possibility that uuid.New()
	// generation function will be called concurrently.
	uuid.EnableRandPool()

	jobs := make([]Job, 0)
	for i := 0; i < 100; i++ {
		id = uuid.New()
		jb = Job{id, []int64{int64(rand.Int63n(100)), rand.Int63n(100)}}
		jobs = append(jobs, jb)
	}
	return jobs
}

func Scheduler(joblist []Job) {

	ch := make(chan jobResponse)
	for _, v := range joblist {
		go worker(ch, v)
	}

	for i := 0; i < len(joblist); i++ {
		//There is a jobs response output
		fmt.Println(<-ch)
		//<-ch
	}

}

func worker(ch chan<- jobResponse, jb Job) error {
	resp := jobResponse{id: jb.id, out: 0}
	for _, v := range jb.data {
		resp.out += v
	}
	runtime.Gosched()
	ch <- resp
	return nil
}
