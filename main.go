package main

import (
	"log"
	"net/http"

	httptransport "github.com/go-kit/kit/transport/http"
)

func main() {
	uss := urlShortenerService{}

	shortenHandler := httptransport.NewServer(
		makeShortenEndpoint(uss),
		decodeUrlShortenerRequest,
		encodeResponse,
	)

	http.Handle("/create", shortenHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
