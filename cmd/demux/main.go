package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"sync"
)

func main() {
	joblist := make([][]int, 0)
	for i := 0; i < 100; i++ {
		joblist = append(joblist, []int{rand.Intn(100), rand.Intn(100)})
	}

	scheduler(joblist)
}

func scheduler(joblist [][]int) {
	//cpu := runtime.NumCPU()
	//fmt.Println("Num CPU: ", cpu)

	wg := sync.WaitGroup{}

	for _, v := range joblist {
		wg.Add(1)
		go worker(&wg, v)
	}
	fmt.Println("Num Goroutine: ", runtime.NumGoroutine())

	wg.Wait()
	fmt.Println("Num Goroutine: ", runtime.NumGoroutine())
}

func worker(w *sync.WaitGroup, job []int) error {
	sum := 0
	for _, v := range job {
		sum += v
	}
	runtime.Gosched()
	fmt.Println("Job: ", job, "\tSUM: ", sum)
	w.Done()
	return nil
}
