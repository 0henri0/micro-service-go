package server

import (
	"log"
	"net/http"
)

type HttpServer struct {
	HttpServer *http.Server
}

func NewApp() (*HttpServer, error) {
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", fs)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		_, err := w.Write([]byte("Hello, World!22"))
		if err != nil {
			log.Fatal(err)
		}
	})

	server := &http.Server{
		Addr: ":5000",
	}

	return &HttpServer{HttpServer: server}, nil
}

func (a *HttpServer) Run() error {
	log.Println("Server is running on port 5000")

	return a.HttpServer.ListenAndServe()
}
