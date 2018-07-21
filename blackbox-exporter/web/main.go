package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/unrolled/render"
)

var (
	renderer = render.New()
)

func probe(w http.ResponseWriter, r *http.Request) {
	renderer.Text(w, http.StatusOK, "")
}

func main() {
	router := chi.NewRouter()

	router.Get("/probe", probe)

	log.Fatal(http.ListenAndServe(":2112", router))
}
