package renderer

import (
	"ego/pkg/configuration"
	"ego/pkg/renderable"
	"fmt"
	"log"
	"net/http"
)

func init() {
	RegisterRendererFactory("net", func(config configuration.Renderer) Renderer {
		return &NetRenderer{}
	})
}

type NetRenderer struct {
}

// IsAsync implements Renderer
func (*NetRenderer) IsAsync() bool {
	return false
}

// Refresh implements Renderer
func (*NetRenderer) Refresh() {

}

// Render implements Renderer
func (*NetRenderer) Render(renderable.Renderable) {

}

// Start implements Renderer
func (*NetRenderer) Start(chan bool) {
	addr := ":8081"
	fmt.Printf("Starting server at address %s \n", addr)
	if err := http.ListenAndServe(addr, nil); err != nil {
		log.Fatal(err)
	}
}

func (r *NetRenderer) Init() {
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/status", statusHandler)
	http.HandleFunc("/render", renderHandler)

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
