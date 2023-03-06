package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/eastata/demux/internal/logger"
	"github.com/eastata/demux/pkg/demux"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"time"
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

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/job_submit", JobSubmit).Methods("POST")

	// Swagger UI from filesystem
	// This will serve files under http://localhost:8000/swaggerui/<filename>
	var dir string
	flag.StringVar(&dir, "dir", "./swaggerui/", "the directory to serve files from. Defaults to the current dir")
	flag.Parse()

	router.PathPrefix("").Handler(http.StripPrefix("/swaggerui/", http.FileServer(http.Dir("./swaggerui"))))

	address := "0.0.0.0"
	port := "8080"

	srv := &http.Server{
		Handler:      router,
		Addr:         fmt.Sprintf("%s:%s", address, port),
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	logger.Info(fmt.Sprintf("Starting Demux API-Server on %s:%s ...", address, port))
	log.Fatal(srv.ListenAndServe())

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
