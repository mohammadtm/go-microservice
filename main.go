package main

import (
	"log"
	"net/http"
)

func main() {
	server := &http.Server{
		Addr: ":8080",
		Handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
			w.Write([]byte("Hello World"))
		}),
	}
	err := server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
