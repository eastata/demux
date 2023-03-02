package main

import (
	"encoding/json"
	"flag"
	"github.com/eastata/demux/pkg/demux"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"time"
)

// swagger:model CommonError
type CommonError struct {
	// Status of the error
	// in: int64
	Status int64 `json:"status"`
	// Message of the error
	// in: string
	Message string `json:"message"`
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

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/job_submit", JobSubmit).Methods("POST")

	// Swagger UI from filesystem
	// This will serve files under http://localhost:8000/swaggerui/<filename>
	var dir string
	flag.StringVar(&dir, "dir", "./swaggerui/", "the directory to serve files from. Defaults to the current dir")
	flag.Parse()

	router.PathPrefix("").Handler(http.StripPrefix("/swaggerui/", http.FileServer(http.Dir("./swaggerui"))))

	srv := &http.Server{
		Handler:      router,
		Addr:         "127.0.0.1:8080",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())

}

// JobSubmit serves the submitting jobs
func JobSubmit(w http.ResponseWriter, r *http.Request) {

	// swagger:route POST /job_submit JobSubmit
	//
	// Submit the job for summing the list of int64
	//
	// This will submit the job to demux and return JobID
	//
	//     Responses:
	//       401: CommonError

	w.Header().Set("Content-Type", "application/json")
	var v jobRequest
	err := json.NewDecoder(r.Body).Decode(&v)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	jb := demux.NewJob(v.Data)
	demux.Scheduler([]demux.Job{jb})

	json.NewEncoder(w).Encode(v)

}
