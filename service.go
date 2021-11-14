package main

import (
	"github.com/segmentio/ksuid"
)

type UrlShortenerService interface {
	Shorten(string) (string, error)
}

type urlShortenerService struct{}

func (urlShortenerService) Shorten(url string) (surl string, err error) {
	// Return uuid as short URL
	surl = ksuid.New().String()
	return
}
