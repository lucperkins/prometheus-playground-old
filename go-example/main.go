package main

import (
	"github.com/lucperkins/prometheus-playground/go/web"

	log "github.com/sirupsen/logrus"
)

func main() {
	log.SetFormatter(&log.JSONFormatter{})

	httpServer := web.NewHttpServer()

	log.Fatal(httpServer.Start())
}
