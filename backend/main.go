package main

import (
	"log"
	"net/http"
	"os"
)

func main() {
	hostname, _ := os.Hostname()
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Printf("[backend] %s, request from: %s", hostname, r.RemoteAddr)

		resp := "---------------------------\n"
		resp += "backend " + hostname + "\n"
		resp += "---------------------------\n"
		for _, env := range os.Environ() {
			resp += env + "\n"
		}
		w.Write([]byte(resp))
	})
	http.ListenAndServe(":8080", nil)
}
