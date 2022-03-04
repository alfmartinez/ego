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
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/status", statusHandler)
	http.HandleFunc("/render", renderHandler)

	fmt.Printf("Starting server at address %s \n", addr)
	if err := http.ListenAndServe(addr, nil); err != nil {
		log.Fatal(err)
	}
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Index")
}

func statusHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Status")
}
func renderHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Render")
}
