package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	// Connect to database, or Crash
	if err := connectDatabase(); err != nil {
		log.Fatal(err)
	}
	defer databaseConn.Close()
	
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "hello gopher")
	})

	http.HandleFunc("/aligator", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hi Mr Aligator")
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}