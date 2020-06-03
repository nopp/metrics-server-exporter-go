package node

import (
	"encoding/json"
	"metrics-server-exporter-go/api"
	"regexp"
	"strconv"
	"time"

	"github.com/prometheus/client_golang/prometheus"
)

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
	MetricsNodesCPU = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "kube_metrics_server_nodes_cpu",
			Help: "Metrics Server Nodes CPU",
		},
		[]string{"instance"},
	)
	MetricsNodesMEM = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "kube_metrics_server_nodes_mem",
			Help: "Metrics Server Nodes Memory",
		},
		[]string{"instance"},
	)
)

func Collect() {

	re := regexp.MustCompile("[^0-9]")

	// NODE
	apiNode := api.Connect("node")

	var nodes Info
	_ = json.NewDecoder(apiNode.Body).Decode(&nodes)

	for i := range nodes.Items {

		// Only numbers :D
		nodes.Items[i].Usage.CPU = re.ReplaceAllLiteralString(nodes.Items[i].Usage.CPU, "")

		// String to float
		CPUfloat, _ := strconv.ParseFloat(nodes.Items[i].Usage.CPU, 64)

		// Only numbers :D
		nodes.Items[i].Usage.Memory = re.ReplaceAllLiteralString(nodes.Items[i].Usage.Memory, "")

		// String to float
		MEMfloat, _ := strconv.ParseFloat(nodes.Items[i].Usage.Memory, 64)

		MetricsNodesCPU.WithLabelValues(nodes.Items[i].Metadata.Name).Add(CPUfloat)
		MetricsNodesMEM.WithLabelValues(nodes.Items[i].Metadata.Name).Add(MEMfloat)
	}
}
