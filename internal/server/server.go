package server

import (
	"fmt"
	"log"
	"net/http"
)

type Server interface {
	Start(addr string)
}

func CreateServer() Server {
	return &server{}
}

type server struct{}

func (s *server) Start(addr string) {
	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello!")
	})

	fmt.Printf("Starting server at address %s \n", addr)
	if err := http.ListenAndServe(addr, nil); err != nil {
		log.Fatal(err)
	}
}
