package api

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/skiba-mateusz/task-manager/services/user"
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

	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	userStore := user.NewStore(s.db)
	
	userHandler := user.NewHandler(userStore)

	r.Route("/v1", func(r chi.Router) {
		r.Route("/users", func(r chi.Router) {
			r.Post("/register", userHandler.RegisterUser)
		})
	})

	log.Println("server is listening on", s.addr)

	return http.ListenAndServe(s.addr, r)
}