package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

const (
	PORT = 2112
)

var (
	queuedOps = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "myapp_ops_queued",
		Help: "Something or other",
	})
)

func recordMetrics() {
	go func() {
		for {
			queuedOps.Add(1)
			time.Sleep(2 * time.Second)
		}
	}()
}

func init() {
	prometheus.MustRegister(queuedOps)
}

func main() {
	address := fmt.Sprintf(":%d", PORT)

	recordMetrics()

	http.Handle("/metrics", promhttp.Handler())
	log.Fatal(http.ListenAndServe(address, nil))
}
