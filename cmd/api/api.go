package api

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
)

type APIServer struct {
	addr 	string
	db		*sql.DB
}


func NewApiServer(addr string, db *sql.DB) *APIServer {
	return &APIServer{
		addr: addr,
		db: db,
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