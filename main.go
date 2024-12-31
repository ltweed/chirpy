package main

import (
	"log"
	"net/http"
)

func main() {
	const port = "8080"

	mux := http.NewServeMux()

	// Use http.FileServer to serve files from the current directory
	fileServer := http.FileServer(http.Dir("."))

	// Handle the root path ("/") and serve the index.html file
	mux.Handle("/", fileServer)
	
	srv := &http.Server{
		Addr:    ":" + port,
		Handler: mux,
	}

	log.Printf("Serving on port: %s\n", port)
	log.Fatal(srv.ListenAndServe())
}

