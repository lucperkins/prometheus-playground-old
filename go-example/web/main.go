package main

import (
	"fmt"
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

func (h *handler) greeting(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")

	if name == "" {
		h.r.Text(w, http.StatusBadRequest, "You must supply a name as a query parameter")
	}

	h.r.Text(w, http.StatusOK, fmt.Sprintf("Hello, %s!", name))
}

func main() {
	r := chi.NewRouter()

	h := newHandler()

	r.Get("/greeting", h.greeting)
	r.Handle("/metrics", promhttp.Handler())

	log.Print("Starting up the server on port 2112")

	log.Fatal(http.ListenAndServe(":2112", r))
}
