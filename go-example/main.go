package main

import (
	"net/http"
	"runtime"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	numCpus = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "num_cpus",
		Help: "The number of CPUs on this machine",
	})
)

func main() {
	prometheus.MustRegister(numCpus)

	router := http.NewServeMux()

	numCpus.Set(float64(runtime.NumCPU()))

	router.Handle("/metrics", promhttp.Handler())

	server := &http.Server{
		Addr:    ":2112",
		Handler: router,
	}

	server.ListenAndServe()
}
