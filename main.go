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
	addr := "localhost:8080"
	if len(os.Args) > 1 {
		addr := os.Args[1]
		if addr == "" {
			addr = ":8000"
		}
	}
	log.Println("Server running at " + addr)
	log.Fatal(http.ListenAndServe(addr, Logger(http.FileServer(http.Dir(d)))))
}

// Logger middleware for HTTP handling.
func Logger(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("[%s] [%s] %s", r.RemoteAddr, r.Method, r.URL.String())
		h.ServeHTTP(w, r)
	})
}
