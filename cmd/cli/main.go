package main

import (
	"github.com/eastata/demux/pkg/demux"
)

func main() {
	jobslist := demux.JobsGenerator()
	demux.Scheduler(jobslist)
}
