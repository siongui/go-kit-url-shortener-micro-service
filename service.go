package main

type UrlShortenerService interface {
	Shorten(string) (string, error)
}

type urlShortenerService struct{}

func (urlShortenerService) Shorten(url string) (surl string, err error) {
	// To be implemented later. Just return original URL
	return url, nil
}
