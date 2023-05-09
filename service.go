package main

import (
	"context"
	"encoding/json"
	"net/http"
)

type Service interface {
	GetCatFact(context.Context) (*GetCatFactResponse, error) 
}

type GetCatFactService struct {
	url string
}

func NewGetCatFactService(url string) Service {
	return &GetCatFactService{url: url}
}

func (s *GetCatFactService) GetCatFact(ctx context.Context) (*GetCatFactResponse, error){
	r, e := http.Get(s.url)
	if e != nil {
		return nil, e
	}
	defer r.Body.Close()
	fact := &GetCatFactResponse{}
	if e := json.NewDecoder(r.Body).Decode(fact); e != nil {
		return nil, e
	}
	return fact, nil
	
}