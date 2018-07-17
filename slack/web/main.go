package main

import (
	"net/http"

	"github.com/prometheus/client_golang/prometheus/promhttp"

	"github.com/go-chi/chi"
	"github.com/unrolled/render"
)

func main() {
	renderer := render.New()
	router := chi.NewRouter()

	router.Get("/hello", func(w http.ResponseWriter, r *http.Request) {
		renderer.Text(w, http.StatusOK, "Hello world")
	})

	router.Handle("/metrics", promhttp.Handler())

	http.ListenAndServe(":2112", router)
}
