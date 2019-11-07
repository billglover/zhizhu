package main

import (
	"fmt"
	"net/http"

	"github.com/billglover/zhizhu/model"
)

func (s *server) HandleFeedPost() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		rData := model.FeedInfo{}
		err := s.decode(w, r, &rData)
		if err != nil {
			fmt.Println(err)
			s.respondErr(w, r, err, http.StatusBadRequest)
			return
		}

		f, err := s.fm.Create(rData.Link)
		if err != nil {
			fmt.Println(err)
			s.respondErr(w, r, err, http.StatusInternalServerError)
			return
		}

		s.respond(w, r, f.ID, http.StatusAccepted)
	}
}
