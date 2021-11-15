package main

import (
	"fmt"
	"time"

	"github.com/go-kit/kit/metrics"
)

type instrumentingMiddleware struct {
	requestCount   metrics.Counter
	requestLatency metrics.Histogram
	countResult    metrics.Histogram
	next           UrlShortenerService
}

func (mw instrumentingMiddleware) Shorten(url string) (surl string, err error) {
	defer func(begin time.Time) {
		lvs := []string{"method", "shorten", "error", fmt.Sprint(err != nil)}
		mw.requestCount.With(lvs...).Add(1)
		mw.requestLatency.With(lvs...).Observe(time.Since(begin).Seconds())
	}(time.Now())

	surl, err = mw.next.Shorten(url)
	return
}

func (mw instrumentingMiddleware) GetOriginalUrl(shortcode string) (url string, err error) {
	defer func(begin time.Time) {
		lvs := []string{"method", "get_original_url", "error", fmt.Sprint(err != nil)}
		mw.requestCount.With(lvs...).Add(1)
		mw.requestLatency.With(lvs...).Observe(time.Since(begin).Seconds())
	}(time.Now())

	url, err = mw.next.GetOriginalUrl(shortcode)
	return
}
