package web

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/prometheus/client_golang/prometheus"
	log "github.com/sirupsen/logrus"
)

type HttpServer struct {
	routes *chi.Mux
}

func NewHttpServer() *HttpServer {
	r := chi.NewRouter()

	r.Handle("/metrics", prometheus.Handler())

	return &HttpServer{
		routes: r,
	}
}

func (s *HttpServer) Start() error {
	log.Print("Starting up the server on port 2112")

	return http.ListenAndServe(":2112", s.routes)
}
