package main

func (s *server) routes() {
	s.router.HandleFunc("/feed", s.HandleFeedPost()).Methods("POST")
}
