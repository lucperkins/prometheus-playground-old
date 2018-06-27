package main

import (
	"log"
	"net/http"

	"github.com/unrolled/render"

	"github.com/prometheus/client_golang/prometheus/promhttp"

	"github.com/go-chi/chi"
)

type handler struct {
	r *render.Render
}

func newHandler() *handler {
	return &handler{
		r: render.New(),
	}
}

func (h *handler) get(w http.ResponseWriter, r *http.Request) {
	h.r.Text(w, http.StatusOK, "Hello world!!!")
}

func main() {
	r := chi.NewRouter()

	h := newHandler()

	r.Get("/", h.get)
	r.Handle("/metrics", promhttp.Handler())

	log.Print("Starting up the server on port 2112")

	log.Fatal(http.ListenAndServe(":2112", r))
}
