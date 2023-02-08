package main

import (
	"context"
	"encoding/json"
	"net/http"
)

type ApiServer struct {
	svc Service
}

func NewApiServer(svc Service) *ApiServer {
	return &ApiServer{
		svc: svc,
	}
}

func (s *ApiServer) GetRandomChuckNorrisJokesHandler(w http.ResponseWriter, r *http.Request) {
	jokes, err := s.svc.GetRandomChuckNorrisJokes(context.Background())
	if err != nil {
		response := &ResponseApi{
			Message: err.Error(),
			Data:    nil,
			Meta:    nil,
			Status:  false,
		}
		WriteJSON(w, http.StatusUnprocessableEntity, response)
		return
	}

	response := &ResponseApi{
		Message: "success get jokes data",
		Data:    jokes,
		Meta:    nil,
		Status:  true,
	}
	WriteJSON(w, http.StatusOK, response)
}

func WriteJSON(w http.ResponseWriter, s int, v any) error {
	w.WriteHeader(s)
	w.Header().Add("Content-Type", "application/json")
	return json.NewEncoder(w).Encode(v)
}
