package main

import (
	"encoding/json"
	"fmt"
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

// swagger:model JobSubmit
type jobRequest struct {
	// Slice of int to sum
	// in: int64
	Job []int64 `json:"job" validate:"required,gt=1,drive,numeric"`
}

func main() {
	router := mux.NewRouter()

	// Swagger UI from filesystem
	fs := http.FileServer(http.Dir("./swaggerui"))
	router.Handle("/swaggerui/{rest}", http.StripPrefix("/swaggerui", fs))

	router.HandleFunc("/", JobSubmit).Methods("POST")

	srv := &http.Server{
		Handler:      router,
		Addr:         "127.0.0.1:8080",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())

}

// swagger:route POST / admin JobSubmit
// Submit job
//
// security:
// - apiKey: []
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

	fmt.Println(v)
	fmt.Fprintf(w, "Body: %+v", v)

}
