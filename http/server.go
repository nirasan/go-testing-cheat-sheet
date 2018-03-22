package http

import "net/http"

type MyServer struct{}

func (s *MyServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	m := http.NewServeMux()
	m.HandleFunc("/hello", HelloHandler)
	m.HandleFunc("/json", JsonHandler)
	m.ServeHTTP(w, r)
}
