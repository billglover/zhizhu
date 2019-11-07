package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/billglover/zhizhu/feed"
	"github.com/billglover/zhizhu/log"
	"github.com/gorilla/mux"
)

type server struct {
	fm     *feed.Manager
	router *mux.Router
	log    log.Logger
}

func newServer() *server {
	var s = new(server)
	s.router = mux.NewRouter()
	s.routes()
	return s
}

func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}

func (s *server) respond(w http.ResponseWriter, r *http.Request, data interface{}, status int) {
	w.WriteHeader(status)
	if data != nil {
		err := json.NewEncoder(w).Encode(data)
		if err != nil {
			s.log.Info(err)
		}
	}
}

func (s *server) respondErr(w http.ResponseWriter, r *http.Request, err error, status int) {
	type response struct {
		Status  int    `json:"status,omitempty"`
		Message string `json:"message,omitempty"`
	}

	data := response{Status: status, Message: fmt.Sprintf("%s: %s", http.StatusText(status), err.Error())}

	w.WriteHeader(status)
	err = json.NewEncoder(w).Encode(data)
	if err != nil {
		s.log.Info(err)
	}
}

func (s *server) decode(w http.ResponseWriter, r *http.Request, v interface{}) error {
	return json.NewDecoder(r.Body).Decode(v)
}
