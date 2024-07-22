package api

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
)

type APIServer struct {
	addr string
}


func NewApiServer(addr string) *APIServer {
	return &APIServer{
		addr: addr,
	}
}

func (s *APIServer) Run() error {
	r := chi.NewRouter()

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("It works"))
	})

	log.Println("server is listening on", s.addr)

	return http.ListenAndServe(s.addr, r)
}