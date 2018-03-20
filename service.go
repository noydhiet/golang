package main

import ("sync"
	"context"
	"errors"
	"strings")

// Counter adds a value and returns a new value
type Counter interface {
	Add(int) int
	Uppercase(context.Context, string) (string, error)
}

type countService struct {
	v  int
	mu sync.Mutex
}

func (countService) Uppercase(_ context.Context, s string) (string, error) {
	if s == "" {
		return "", ErrEmpty
	}
	return strings.ToUpper(s), nil
}

func (c *countService) Add(v int) int {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.v += v
	return c.v
}

var ErrEmpty = errors.New("Empty string")
