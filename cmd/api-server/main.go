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

// swagger:parameters admin JobSubmit
type jobRequest struct {
	// Slice of int64 to sum
	// in: body
	// schema:
	// 	type: object
	// 	required:
	// 		- body
	// 	properties: {body: {type: array, items: [int64]}}
	Job []int64 `json:"job" validate:"required,gt=1,drive,numeric"`
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

// swagger:route POST /job_submit admin JobSubmit
// Submit the job of summing the list of int64
//
// security:
// - apiKey: []
//
// responses:
//
//	401: CommonError
func JobSubmit(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var v jobRequest
	err := json.NewDecoder(r.Body).Decode(&v)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	jb := demux.NewJob(v.Job)
	demux.Scheduler([]demux.Job{jb})

	json.NewEncoder(w).Encode(v)

}
