package main

import (
	"log"
	"net/http"

	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/unrolled/render"

	"github.com/go-chi/chi"
)

type httpServer struct {
	r *render.Render
}

func (s *httpServer) get(w http.ResponseWriter, r *http.Request) {
	s.r.Text(w, http.StatusOK, "Hello world!")
}

func main() {
	r := chi.NewRouter()

	server := httpServer{
		r: render.New(),
	}

	r.Get("/", server.get)
	r.Handle("/metrics", promhttp.Handler())

	log.Print("Starting up the server on port 9999")

	log.Fatal(http.ListenAndServe(":9999", r))
}
