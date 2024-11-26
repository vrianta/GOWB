package Server

import (
	"fmt"
	Application "gowb/app"
	"log"
	"net/http"
)

type Server struct{}

func (s *Server) RunServer() {
	// Register the handler function for the root URL
	http.HandleFunc("/", s.handler)

	// Start the server on port 8080
	server_config, err := Application.LoadConfig("server.config")
	if err != nil {
		log.Fatalln(err)
	}
	// log.Println("Server is listening on port " + app.server_config["port"] + "...")
	if err := http.ListenAndServe(server_config["host"]+":"+server_config["port"], nil); err != nil {
		log.Fatal(err)
	}
}

// Handler function for the root URL
func (s *Server) handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, World!")
	// r.URL.
}
