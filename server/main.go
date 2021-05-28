package main

import (
	"flag"
	"log"
	"net/http"
)

func main() {
	var (
		resp string
	)

	flag.StringVar(&resp, "resp", "", "response value")
	flag.Parse()

	http.HandleFunc("/healthz", func(rw http.ResponseWriter, r *http.Request) {
		rw.WriteHeader(200)
		return
	})
	http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		rw.Write([]byte(resp))
		return
	})

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
