package main

import (
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/unrolled/render"
)

const (
	serviceName = "web"
)

var (
	prometheusLabels = prometheus.Labels{"service": serviceName}
)

type server struct {
	httpReqsCounter *prometheus.CounterVec
	latencyHist     *prometheus.HistogramVec
	renderer        *render.Render
}

func newServer() *server {
	requestInfoCounter := prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name:        "http_request_info",
			Help:        "HTTP request counter by response code, request method, and request path",
			ConstLabels: prometheusLabels,
		},
		[]string{"code", "method", "path"},
	)

	prometheus.MustRegister(requestInfoCounter)

	latencyHistogram := prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Name:        "http_latency_milliseconds",
			Help:        "Per-request latency in milliseconds",
			ConstLabels: prometheusLabels,
		},
		[]string{"code", "method", "path"},
	)

	prometheus.MustRegister(latencyHistogram)

	return &server{
		httpReqsCounter: requestInfoCounter,
		latencyHist:     latencyHistogram,
		renderer:        render.New(),
	}
}

func (s *server) prometheusMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		ww := middleware.NewWrapResponseWriter(w, r.ProtoMajor)

		next.ServeHTTP(w, r)

		code := http.StatusText(ww.Status())
		method := strings.ToLower(r.Method)
		path := r.URL.Path

		if r.RequestURI != "/metrics" {
			s.httpReqsCounter.WithLabelValues(
				code,
				method,
				path,
			).Inc()

			requestLatency := float64(time.Since(start).Nanoseconds())

			s.latencyHist.WithLabelValues(
				code,
				method,
				path,
			).Observe(requestLatency)
		}
	})
}

func (s *server) get(w http.ResponseWriter, r *http.Request) {
	s.renderer.JSON(w, http.StatusOK, map[string]string{"hello": "world"})
}

func (s *server) wrong(w http.ResponseWriter, r *http.Request) {
	s.renderer.Text(w, http.StatusInternalServerError, "Something went wrong")
}

func main() {
	router := chi.NewRouter()
	server := newServer()

	router.Use(server.prometheusMiddleware)

	router.Get("/", server.get)

	router.Get("/wrong", server.wrong)

	router.Handle("/metrics", promhttp.Handler())

	log.Fatal(http.ListenAndServe(":2112", router))
}
