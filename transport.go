package main

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/go-kit/kit/endpoint"
)

type urlShortenerRequest struct {
	Url string `json:"url"`
}

type urlShortenerResponse struct {
	ShortUrl string `json:"short_url"`
	Err      string `json:"err,omitempty"` // errors don't JSON-marshal, so we use a string
}

func makeShortenEndpoint(uss UrlShortenerService) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		req := request.(urlShortenerRequest)
		v, err := uss.Shorten(req.Url)
		if err != nil {
			return urlShortenerResponse{v, err.Error()}, nil
		}
		return urlShortenerResponse{v, ""}, nil
	}
}

func decodeUrlShortenerRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request urlShortenerRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}

func encodeResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}
