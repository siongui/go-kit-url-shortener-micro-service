package main

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/go-kit/kit/log"
	kitprometheus "github.com/go-kit/kit/metrics/prometheus"
	httptransport "github.com/go-kit/kit/transport/http"
	stdprometheus "github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func makeHttpHandler(uss UrlShortenerService) http.Handler {
	r := gin.Default()

	shortenHandler := httptransport.NewServer(
		makeShortenEndpoint(uss),
		decodeUrlShortenerRequest,
		encodeResponse,
	)
	r.POST("/create", gin.WrapH(shortenHandler))

	r.GET("/metrics", gin.WrapH(promhttp.Handler()))

	getOriginalUrlHandler := httptransport.NewServer(
		makeGetOriginalUrlEndpoint(uss),
		decodeUrlGetOriginalUrlRequest,
		encodeResponse,
	)
	r.GET("/:shortcode", gin.WrapH(getOriginalUrlHandler))

	return r
}

func main() {
	logger := log.NewLogfmtLogger(os.Stderr)

	fieldKeys := []string{"method", "error"}
	requestCount := kitprometheus.NewCounterFrom(stdprometheus.CounterOpts{
		Namespace: "my_group",
		Subsystem: "url_shortener_service",
		Name:      "request_count",
		Help:      "Number of requests received.",
	}, fieldKeys)

	requestLatency := kitprometheus.NewSummaryFrom(stdprometheus.SummaryOpts{
		Namespace: "my_group",
		Subsystem: "url_shortener_service",
		Name:      "request_latency_microseconds",
		Help:      "Total duration of requests in microseconds.",
	}, fieldKeys)

	// The following countResult is not used. Can be removed.
	countResult := kitprometheus.NewSummaryFrom(stdprometheus.SummaryOpts{
		Namespace: "my_group",
		Subsystem: "url_shortener_service",
		Name:      "count_result",
		Help:      "The result of each count method.",
	}, []string{}) // no fields here

	var uss UrlShortenerService
	uss = urlShortenerService{ds: getSqliteDataSource()}
	uss = loggingMiddleware{logger, uss}
	uss = instrumentingMiddleware{requestCount, requestLatency, countResult, uss}

	logger.Log("msg", "HTTP", "addr", ":8080")
	logger.Log("err", http.ListenAndServe(":8080", makeHttpHandler(uss)))
}

func getSqliteDataSource() DataSource {
	// Used by services to store short URL
	s3 := sqlite{}
	err := s3.InitSQLite(false)
	if err != nil {
		panic(err)
	}
	_, err = s3.CreateShortUrlTable()
	if err != nil {
		panic(err)
	}

	return &s3
}
