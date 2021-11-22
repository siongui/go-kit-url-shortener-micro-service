package main

import (
	"database/sql"

	"github.com/matoous/go-nanoid/v2"
	"github.com/siongui/go-kit-url-shortener-micro-service/datasource"
)

// https://en.wikipedia.org/wiki/Base62
var base58 = "123456789ABCDEFGHJKLMNPQRSTUVWXYZabcdefghijkmnopqrstuvwxyz"

type UrlShortenerService interface {
	Shorten(string) (string, error)
	GetOriginalUrl(string) (string, error)
}

type urlShortenerService struct {
	ds datasource.DataSource
}

func (u urlShortenerService) createShortUrlCode(url string) (shortcode string, err error) {
	shortcode, err = gonanoid.Generate(base58, 5)
	if err != nil {
		return
	}

	_, err = u.ds.SelectByShortUrlCode(shortcode)
	if err == sql.ErrNoRows {
		return shortcode, nil
	}
	if err == nil {
		return u.createShortUrlCode(url)
	}
	return
}

func (u urlShortenerService) Shorten(url string) (surl string, err error) {
	su, err := u.ds.SelectByOriginalUrl(url)
	if err == sql.ErrNoRows {
		surl, err := u.createShortUrlCode(url)
		if err == nil {
			_, err = u.ds.InsertShortUrl(datasource.ShortUrl{surl, url})
		}
		return surl, err
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
