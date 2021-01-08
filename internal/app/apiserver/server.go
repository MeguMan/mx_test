package apiserver

import (
	"encoding/json"
	"fmt"
	"github.com/MeguMan/mx_test/internal/app/model"
	"github.com/MeguMan/mx_test/internal/app/store/postgres_store"
	"github.com/MeguMan/mx_test/internal/app/xlsxDecoder"
	"github.com/gorilla/mux"
	"net/http"
)

type ReqBody struct {
	SellerId int `json:"seller_id"`
	OfferId int `json:"offer_id"`
	Path string `json:"path"`
	Pattern string `json:"pattern"`
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
		rowsStats := model.RowsStats{}
		or := s.store.Offer()
		w.Header().Set("Content-Type", "application/json")
		err := json.NewDecoder(r.Body).Decode(&rb)
		oo, err := xlsxDecoder.ParseFile(rb.Path, &rowsStats)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		for _, o := range oo {
			if o.Available {
				if or.Exists(o.OfferId, o.SellerId) {
					or.Update(&o, &rowsStats)
				} else {
					or.Create(&o, &rowsStats)
				}
			} else {
				or.Delete(&o, &rowsStats)
			}
		}

		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		w.WriteHeader(http.StatusCreated)
		resp, _ := json.Marshal(rowsStats)
		fmt.Fprint(w, string(resp))
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
		oo, _ := s.store.Offer().GetByPattern(rb.OfferId, rb.SellerId, rb.Pattern)
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, oo)
	}
}