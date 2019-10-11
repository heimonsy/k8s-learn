package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func main() {
	hostname, _ := os.Hostname()
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Printf("[frontend] request from: %s", r.RemoteAddr)

		geted, err := http.Get("http://backend-service:18080/")
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}
		defer geted.Body.Close()
		data, err := ioutil.ReadAll(geted.Body)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}

		resp := "---------------------------\n"
		resp += "frontend " + hostname + "\n"
		resp += "---------------------------\n"
		for _, env := range os.Environ() {
			resp += env + "\n"
		}
		w.Write([]byte(resp))
		w.Write([]byte("response from backend >>>>>>>>>>>\n"))
		w.Write(data)

	})
	http.ListenAndServe(":8080", nil)
}
