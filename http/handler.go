package http

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello %s", r.FormValue("name"))
}

type JsonRequest struct {
	Name string
}

type JsonResponse struct {
	Message string
}

func JsonHandler(w http.ResponseWriter, r *http.Request) {
	body := JsonRequest{}
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	b, err := json.Marshal(JsonResponse{Message: "hello " + body.Name})
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprint(w, string(b))
}
