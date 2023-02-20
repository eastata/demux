package main

import (
	"github.com/google/uuid"
	"math/rand"
	"runtime"
)

type job struct {
	id   uuid.UUID
	data []int
}

type jobResponse struct {
	id  uuid.UUID
	out int
}

func main() {
	jobslist := jobsGenerator()
	scheduler(jobslist)
}

func jobsGenerator() []job {
	var (
		jb job
		id uuid.UUID
	)

	// uuid.EnableRandPool() are not thread-safe and should only be called when there is no possibility that uuid.New()
	// generation function will be called concurrently.
	uuid.EnableRandPool()

	jobs := make([]job, 0)
	for i := 0; i < 100; i++ {
		id = uuid.New()
		jb = job{id, []int{rand.Intn(100), rand.Intn(100)}}
		jobs = append(jobs, jb)
	}
	return jobs
}

func scheduler(joblist []job) {
	//cpu := runtime.NumCPU()
	//fmt.Println("Num CPU: ", cpu)

	ch := make(chan jobResponse)
	for _, v := range joblist {
		go worker(ch, v)
	}

	for i := 0; i < len(joblist); i++ {
		//There is a jobs response output
		//fmt.Println(<-ch)
		<-ch
	}

}

func worker(ch chan<- jobResponse, jb job) error {
	resp := jobResponse{id: jb.id, out: 0}
	for _, v := range jb.data {
		resp.out += v
	}
	runtime.Gosched()
	ch <- resp
	return nil
}
