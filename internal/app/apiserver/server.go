package apiserver

import (
	"encoding/json"
	"fmt"
	"github.com/MeguMan/mx_test/internal/app/cache"
	"github.com/MeguMan/mx_test/internal/app/model"
	"github.com/MeguMan/mx_test/internal/app/store"
	"github.com/MeguMan/mx_test/internal/app/store/postgres_store"
	"github.com/MeguMan/mx_test/internal/app/xlsxDecoder"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

type server struct {
	router *mux.Router
	cache  *cache.LRU
	store  postgres_store.Store
	oauthToken string
}

func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}

func NewServer(store postgres_store.Store, token string) *server {
	s := &server{
		router: mux.NewRouter(),
		cache: cache.NewLru(),
		store:  store,
		oauthToken: token,
	}
	s.configureRouter()
	return s
}

func (s *server) configureRouter() {
	s.router.HandleFunc("/offers", s.HandleOffersPost()).Methods("POST")
	s.router.HandleFunc("/offers", s.HandleOffersGet()).Methods("GET")
	s.router.HandleFunc("/offers/status/{id}", s.HandleOffersStatus()).Methods("GET")
}

func (s *server) HandleOffersPost() func(w http.ResponseWriter, r *http.Request) {
	type ReqBody struct {
		SellerId int `json:"seller_id"`
		Path string `json:"path"`
	}

	type RespBody struct {
		Key string `json:"key"`
	}
	return func(w http.ResponseWriter, r *http.Request) {
		rb := ReqBody{}
		or := s.store.Offer()
		rs := &model.RowsStats{}
		g := model.GoroutineStatus{
			RowsStats: rs,
		}
		w.Header().Set("Content-Type", "application/json")
		err := json.NewDecoder(r.Body).Decode(&rb)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		uuidWithHyphen := uuid.New()
		g.Key = strings.Replace(uuidWithHyphen.String(), "-", "", -1)
		go s.decodeAndSave(or,rb.Path, rb.SellerId, g.Key, &g)
		w.WriteHeader(http.StatusCreated)
		resp, _ := json.Marshal(RespBody{Key: g.Key})
		fmt.Fprint(w, string(resp))
	}
}

func (s *server) HandleOffersGet() func(w http.ResponseWriter, r *http.Request) {
	type ReqBody struct {
		SellerId   *int `json:"seller_id"`
		OfferId    *int `json:"offer_id"`
		Pattern *string `json:"pattern"`
	}
	return func(w http.ResponseWriter, r *http.Request) {
		rb := ReqBody{}
		w.Header().Set("Content-Type", "application/json")
		b, _ := ioutil.ReadAll(r.Body)
		err := json.Unmarshal(b, &rb)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		oo, _ := s.store.Offer().GetByPattern(rb.OfferId, rb.SellerId, rb.Pattern)
		resp, _ := json.Marshal(oo)
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, string(resp))
	}
}

func (s *server) HandleOffersStatus() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		vars := mux.Vars(r)
		id := vars["id"]
		i, err := s.cache.Get(id)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)
		resp, _ := json.Marshal(i)
		fmt.Fprint(w,string(resp))
	}
}

func (s *server) decodeAndSave(or store.OfferRepository,path string, sellerId int, uuid string, g *model.GoroutineStatus) {
	s.cache.Set(uuid, g)
	url, err := xlsxDecoder.GetURLForDownloading(path, s.oauthToken)
	if url == "" {
		fmt.Println("url is empty")
		return
	}
	if err != nil {
		fmt.Println(err)
		return
	}
	err = xlsxDecoder.DownloadFile(url, uuid)
	if err != nil {
		fmt.Println(err)
		return
	}
	oo, err := xlsxDecoder.ParseFile(g.RowsStats, uuid, sellerId)
	if err != nil {
		fmt.Println(err)
	}
	err = os.Remove(fmt.Sprintf("%s.xlsx", uuid))
	if err != nil {
		fmt.Println(err)
		return
	}
	for _, o := range oo {
		if o.Available {
			if or.Exists(o.OfferId, o.SellerId) {
				err = or.Update(&o, g.RowsStats)
				if err != nil {
					fmt.Println(err)
				}
			} else {
				err = or.Create(&o, g.RowsStats)
				if err != nil {
					fmt.Println(err)
				}
			}
		} else {
			err = or.Delete(&o, g.RowsStats)
			if err != nil {
				fmt.Println(err)
			}
		}
	}
	g.Finished = true
	s.cache.Set(uuid, g)
}