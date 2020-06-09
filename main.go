package main

import (
	"log"
	"metrics-server-exporter-go/config"
	"metrics-server-exporter-go/node"
	"metrics-server-exporter-go/pod"
	"net/http"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {

	http.Handle("/metrics", promhttp.Handler())

	prometheus.MustRegister(pod.MetricsPodsCPU)
	prometheus.MustRegister(pod.MetricsPodsMEM)
	prometheus.MustRegister(node.MetricsNodesCPU)
	prometheus.MustRegister(node.MetricsNodesMEM)

	go func() {
		for {
			pod.Collect()
			node.Collect()
			time.Sleep(time.Second)
		}
	}()

	log.Fatal(http.ListenAndServe(config.Host+":"+config.Port, nil))
}
