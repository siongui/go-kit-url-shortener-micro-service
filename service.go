package main

import (
	"database/sql"

	"github.com/segmentio/ksuid"
	"github.com/siongui/go-kit-url-shortener-micro-service/datasource"
)

type UrlShortenerService interface {
	Shorten(string) (string, error)
	GetOriginalUrl(string) (string, error)
}

type urlShortenerService struct {
	ds datasource.DataSource
}

func (u urlShortenerService) Shorten(url string) (surl string, err error) {
	su, err := u.ds.SelectByOriginalUrl(url)
	if err == sql.ErrNoRows {
		// create an unique id as short URL
		surl = ksuid.New().String()

		_, err = u.ds.InsertShortUrl(datasource.ShortUrl{surl, url})
		return
	}

	if err != nil {
		return
	}

	surl = su.ShortUrlCode
	return
}

func (u urlShortenerService) GetOriginalUrl(shortcode string) (url string, err error) {
	su, err := u.ds.SelectByShortUrlCode(shortcode)
	url = su.OriginalUrl
	return
}
