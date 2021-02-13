package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

// MUST BE SET by go build -ldflags "-X main.version=999"
// like 0.6.14-0-g26fe727 or 0.6.14-2-g9118702-dirty

var version string // do not remove or modify

func main() {
	port := os.Getenv("PORT")
	router := mux.NewRouter()

	if port == "" {
		port = "8081"
	}

	router.Use(logMiddleware)
	router.HandleFunc("/insert", insertHandler).Methods(http.MethodGet)
	router.HandleFunc("/qsort", qsortHigh).Methods(http.MethodGet)
	router.HandleFunc("/qsortm", qsortMiddle).Methods(http.MethodGet)
	router.HandleFunc("/qsort3", qsortMedian).Methods(http.MethodGet)
	router.HandleFunc("/qsorti", qsortInsert).Methods(http.MethodGet)
	router.HandleFunc("/qsortf", qsortFlag).Methods(http.MethodGet)
	router.HandleFunc("/version", showVersion).Methods(http.MethodGet)

	log.Printf("version %s listening on port %s", version, port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}
