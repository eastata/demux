package main

import (
	"encoding/json"
	"github.com/eastata/demux/pkg/demux"
	"github.com/google/uuid"
	"net/http"
)

// swagger:model JobUUID
type JobID struct {
	// in: string
	// example: {'id": "75a9e835-5cd6-4499-bd2a-a066e335b963"}
	Id uuid.UUID `json:"id"`
}

// swagger:parameters JobSubmit
type jobRequest struct {
	// Send a json body in a request with a key "data" that must be a list of int64
	//
	// in: body
	// schema:
	//   type: string
	// example: {"data": [5,2,7]}
	Data []int64 `json:"data" validate:"required,gt=1,drive,numeric"`
}

// JobSubmit serves the submitting jobs
func JobSubmit(w http.ResponseWriter, r *http.Request) {

	// swagger:route POST /job_submit JobSubmit
	//
	// Submit the job for summing the list of int64
	//
	// This will submit the job to demux
	//
	//     Responses:
	//       200: JobUUID

	w.Header().Set("Content-Type", "application/json")
	var v jobRequest
	err := json.NewDecoder(r.Body).Decode(&v)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	jb := demux.NewJob(v.Data)
	go demux.Scheduler([]demux.Job{jb})

	json.NewEncoder(w).Encode(JobID{jb.Id})

}
