package node

import (
	"encoding/json"
	"log"
	"metrics-server-exporter-go/api"
	"time"

	"github.com/prometheus/client_golang/prometheus"
)

// Info - Data structure of node
type Info struct {
	Items []struct {
		Metadata struct {
			Name              string    `json:"name"`
			CreationTimestamp time.Time `json:"creationTimestamp"`
		} `json:"metadata"`
		Usage struct {
			CPU    string `json:"cpu"`
			Memory string `json:"memory"`
		} `json:"usage"`
	} `json:"items"`
}

var (
	// MetricsNodesCPU - CPU Gauge
	MetricsNodesCPU = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "kube_metrics_server_nodes_cpu",
			Help: "Metrics Server Nodes CPU",
		},
		[]string{"instance"},
	)
	// MetricsNodesMEM - Memory Gauge
	MetricsNodesMEM = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "kube_metrics_server_nodes_mem",
			Help: "Metrics Server Nodes Memory",
		},
		[]string{"instance"},
	)
)

// Collect responsible for get CPU and Memory data
func Collect() {

	log.Println("Starting collect NODE data.")
	var nodes Info

	apiNode := api.Connect("node")

	_ = json.NewDecoder(apiNode.Body).Decode(&nodes)

	for i := range nodes.Items {

		MetricsNodesCPU.WithLabelValues(nodes.Items[i].Metadata.Name).Add(api.ReturnFloat(nodes.Items[i].Usage.CPU))
		MetricsNodesMEM.WithLabelValues(nodes.Items[i].Metadata.Name).Add(api.ReturnFloat(nodes.Items[i].Usage.Memory))
	}
	log.Println("NODE data collected.")
}
