package main

import (
	"context"
	"encoding/json"
	"net/http"
)

type Apiserver struct {
	svc Service
}

func NewApiserver(svc Service) *Apiserver {
	return &Apiserver{svc: svc}
}

func (s *Apiserver) Start(listenAddr string) error {
	http.HandleFunc("/", s.handleGetCatFact)
	return http.ListenAndServe(listenAddr, nil)
}

func (s *Apiserver) handleGetCatFact(w http.ResponseWriter, r *http.Request) {
	fact, e := s.svc.GetCatFact(context.Background())
	if e != nil {
		writeJson(w, http.StatusInternalServerError, map[string]string{"error": e.Error()})
		return
	}
	writeJson(w, http.StatusOK, fact)
}

func writeJson(w http.ResponseWriter, s int, v any) error {
	w.WriteHeader(s)
	w.Header().Add("Content-Type", "application/json")
	return json.NewEncoder(w).Encode(v)
}