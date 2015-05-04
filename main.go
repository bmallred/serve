package main

import (
	"log"
	"net/http"
	"path/filepath"
)

func main() {
	d, err := filepath.Abs(".")
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Server running at http://localhost:8000")
	log.Fatal(http.ListenAndServe(":8000", http.FileServer(http.Dir(d))))
}
