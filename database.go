package main

import (
	"database/sql"
)

type ShortUrl struct {
	ShortUrlCode string
	OriginalUrl  string
}

type DataSource interface {
	SelectAllShortUrl() ([]ShortUrl, error)
	InsertShortUrl(ShortUrl) (sql.Result, error)
	SelectByOriginalUrl(string) (ShortUrl, error)
	SelectByShortUrlCode(string) (ShortUrl, error)
}
