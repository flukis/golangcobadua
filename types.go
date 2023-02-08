package main

type ChuckNorrisJokes struct {
	Jokes string `json:"value"`
}

type ResponseApi struct {
	Status  bool   `json:"status"`
	Data    any    `json:"data"`
	Message string `json:"msg"`
	Meta    any    `json:"meta"`
}
