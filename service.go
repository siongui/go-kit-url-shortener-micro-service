package main

import (
	"database/sql"

	"github.com/segmentio/ksuid"
)

type UrlShortenerService interface {
	Shorten(string) (string, error)
}

type urlShortenerService struct {
	ds DataSource
}

func (u urlShortenerService) Shorten(url string) (surl string, err error) {
	su, err := u.ds.SelectByOriginalUrl(url)
	if err == sql.ErrNoRows {
		// create an unique id as short URL
		surl = ksuid.New().String()

		_, err = u.ds.InsertShortUrl(ShortUrl{surl, url})
		return
	}

	if err != nil {
		return
	}

	surl = su.ShortUrlCode
	return
}
