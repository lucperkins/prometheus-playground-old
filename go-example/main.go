package main

import (
	"log"
	"net/http"
	"strings"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/unrolled/render"
)

type server struct {
	httpReqsCounter *prometheus.CounterVec
	renderer        *render.Render
}

func NewServer() *server {
	counter := prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name:        "http_request_info",
			Help:        "HTTP request counter by response code, request method, and request path",
			ConstLabels: prometheus.Labels{"service": "web"},
		},
		[]string{"code", "method", "path"},
	)

	if err := prometheus.Register(counter); err != nil {
		log.Fatalf("Could not register Prometheus counter: %v", err)
	}

	return &server{
		httpReqsCounter: counter,
		renderer:        render.New(),
	}
}

func (s *server) prometheusMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ww := middleware.NewWrapResponseWriter(w, r.ProtoMajor)

		next.ServeHTTP(w, r)

		status := http.StatusText(ww.Status())
		method := strings.ToLower(r.Method)
		path := r.URL.Path

		if r.RequestURI != "/metrics" {
			s.httpReqsCounter.WithLabelValues(
				status,
				method,
				path,
			).Inc()
		}
	})
}

func (s *server) get(w http.ResponseWriter, r *http.Request) {
	s.renderer.JSON(w, http.StatusOK, map[string]string{"foo": "bar"})
}

func main() {
	router := chi.NewRouter()
	server := NewServer()

	router.Use(server.prometheusMiddleware)

	router.Get("/", server.get)

	router.Handle("/metrics", promhttp.Handler())

	log.Fatal(http.ListenAndServe(":2112", router))
}
