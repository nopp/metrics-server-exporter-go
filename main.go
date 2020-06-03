package main

import (
	"net/http"

	"metrics-server-exporter-go/config"
	"metrics-server-exporter-go/prometheus"

	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {

	prometheus.Collect()

	http.Handle("/metrics", promhttp.Handler())
	http.ListenAndServe(config.Host+":"+config.Port, nil)
}
