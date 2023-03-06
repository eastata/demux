package main

import (
	"flag"
	"fmt"
	"github.com/eastata/demux/internal/logger"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	config()

	router := mux.NewRouter()

	router.HandleFunc("/job_submit", JobSubmit).Methods("POST")

	// Swagger UI from filesystem
	// This will serve files under http://{SERVER_ADDRESS}:{SERVER_PORT}/swaggerui/<filename>
	var dir string
	flag.StringVar(&dir, "dir", "./swaggerui/",
		"the directory to serve files from. Defaults to the current dir")
	flag.Parse()

	router.PathPrefix("").Handler(http.StripPrefix("/swaggerui/", http.FileServer(http.Dir("./swaggerui"))))

	address := os.Getenv("SERVER_ADDRESS")
	port := os.Getenv("SERVER_PORT")

	srv := &http.Server{
		Handler:      router,
		Addr:         fmt.Sprintf("%s:%s", address, port),
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	logger.Info(fmt.Sprintf("Starting Demux API-Server on %s:%s ...", address, port))
	log.Fatal(srv.ListenAndServe())
}
