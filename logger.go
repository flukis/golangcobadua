package main

import (
	"context"
	"fmt"
	"time"
)

type LoggingMiddleware struct {
	next Service
}

func NewLoggingMiddleware(svc Service) Service {
	return &LoggingMiddleware{
		next: svc,
	}
}

func (s *LoggingMiddleware) GetRandomChuckNorrisJokes(ctx context.Context) (res *ChuckNorrisJokes, err error) {
	defer func(start time.Time) {
		if err != nil {
			fmt.Printf("error: %s", err.Error())
		}
		fmt.Printf("took\t: %v\njokes\t: %s\n\n", time.Since(start), res.Jokes)
	}(time.Now())

	return s.next.GetRandomChuckNorrisJokes(ctx)
}
