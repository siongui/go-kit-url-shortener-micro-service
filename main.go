package main

import (
	"net/http"
	"os"

	"github.com/go-kit/kit/log"
	httptransport "github.com/go-kit/kit/transport/http"
)

func main() {
	logger := log.NewLogfmtLogger(os.Stderr)

	var uss UrlShortenerService
	uss = urlShortenerService{}
	uss = loggingMiddleware{logger, uss}

	shortenHandler := httptransport.NewServer(
		makeShortenEndpoint(uss),
		decodeUrlShortenerRequest,
		encodeResponse,
	)

	http.Handle("/create", shortenHandler)
	logger.Log("msg", "HTTP", "addr", ":8080")
	logger.Log("err", http.ListenAndServe(":8080", nil))
}
