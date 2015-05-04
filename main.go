package main

import (
	"log"
	"net/http"
	"os"
	"path/filepath"
)

func main() {
	d, err := filepath.Abs(".")
	if err != nil {
		log.Fatal(err)
	}

	addr := os.Args[1]
	if addr == "" {
		addr = ":8000"
	}

	log.Println("Server running at " + addr)
	log.Fatal(http.ListenAndServe(addr, http.FileServer(http.Dir(d))))
}
