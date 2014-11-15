package main

import (
	"log"
	"net/http"
)

func main() {
	log.Println("Starting the odin http server...")

	http.Handle("/", http.HandlerFunc(dispatch))
	err := http.ListenAndServe("127.0.0.1:8080", nil)
	if err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}

func dispatch(w http.ResponseWriter, req *http.Request) {
	log.Println("req:", req.URL)
}
