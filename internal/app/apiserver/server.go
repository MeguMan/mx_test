package apiserver

import (
	"encoding/json"
	"fmt"
	"github.com/MeguMan/mx_test/internal/app/store/postgres_store"
	"github.com/MeguMan/mx_test/internal/app/xlsxDecoder"
	"github.com/gorilla/mux"
	"net/http"
)

type ReqBody struct {
	SellerId int
	OfferId int
	Path string
	Pattern string
}

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
	s.configureRouter()
	return s
}

func (s *server) configureRouter() {
	s.router.HandleFunc("/offers", s.HandleOffersPost()).Methods("POST")
	s.router.HandleFunc("/offers", s.HandleOffersGet()).Methods("GET")
}

func (s *server) HandleOffersPost() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		rb := ReqBody{}
		or := s.store.Offer()
		w.Header().Set("Content-Type", "application/json")
		err := json.NewDecoder(r.Body).Decode(&rb)
		oo := xlsxDecoder.ParseFile(rb.Path)
		for _, o := range oo {
			or.Create(&o)
		}

		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		w.WriteHeader(http.StatusCreated)
	}
}

func (s *server) HandleOffersGet() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		rb := ReqBody{}
		w.Header().Set("Content-Type", "application/json")
		err := json.NewDecoder(r.Body).Decode(&rb)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		fmt.Println(rb)
		s.store.Offer().GetAll(rb.OfferId, rb.SellerId, rb.Pattern)
		w.WriteHeader(http.StatusOK)
	}
}