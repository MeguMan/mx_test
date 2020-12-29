package apiserver

import (
	"github.com/MeguMan/mx_test/internal/app/store/postgres_store"
	"github.com/gorilla/mux"
	"net/http"
)

type server struct {
	router *mux.Router
	store  postgres_store.Store
}

func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}

func NewServer(store postgres_store.Store) *server {
	s := &server{
		router: mux.NewRouter(),
		store:  store,
	}
/*	s.configureRouter()*/
	return s
}