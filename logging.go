package main

import (
	"time"

	"github.com/go-kit/kit/log"
)

type loggingMiddleware struct {
	logger log.Logger
	next   UrlShortenerService
}

func (mw loggingMiddleware) Shorten(url string) (surl string, err error) {
	defer func(begin time.Time) {
		mw.logger.Log(
			"method", "shorten",
			"input", url,
			"output", surl,
			"err", err,
			"took", time.Since(begin),
		)
	}(time.Now())

	surl, err = mw.next.Shorten(url)
	return
}
