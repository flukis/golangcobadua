package main

import (
	"context"
	"encoding/json"
	"net/http"
)

type Service interface {
	GetRandomChuckNorrisJokes(context.Context) (*ChuckNorrisJokes, error)
}

type ChuckNorrisJokesService struct {
	url string
}

func NewChuckNorrisJokesService(url string) Service {
	return &ChuckNorrisJokesService{
		url: url,
	}
}

func (s *ChuckNorrisJokesService) GetRandomChuckNorrisJokes(ctx context.Context) (*ChuckNorrisJokes, error) {
	resp, err := http.Get(s.url)
	if err != nil {
		return nil, err
	}

	jokes := &ChuckNorrisJokes{}
	if err := json.NewDecoder(resp.Body).Decode(jokes); err != nil {
		return nil, err
	}

	return jokes, nil
}
