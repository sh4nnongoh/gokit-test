package main

import (
	"errors"
	"strings"
)

// Services are represented as interfaces
type StringService interface {
	Uppercase(string) (string, error)
	Count(string) int
}

type stringService struct{}

func (stringService) Uppercase(s string) (string, error) {
	if s == "" {
		return "", ErrEmpty
	}
	return strings.ToUpper(s), nil
}

func (stringService) Count(s string) int {
	return len(s)
}

// ErrEmpty is eturned when input string is empty
var ErrEmpty = errors.New("Empty string")
