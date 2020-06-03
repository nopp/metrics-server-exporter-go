package pod

import (
	"encoding/json"
	"log"
	"metrics-server-exporter-go/api"
	"regexp"
	"strconv"
	"time"

	"github.com/prometheus/client_golang/prometheus"
)

// Info - Data structure of pod
type Info struct {
	Items []struct {
		Metadata struct {
			Name              string    `json:"name"`
			Namespace         string    `json:"namespace"`
			CreationTimestamp time.Time `json:"creationTimestamp"`
		} `json:"metadata"`
		Containers []struct {
			Name  string `json:"name"`
			Usage struct {
				CPU    string `json:"cpu"`
				Memory string `json:"memory"`
			} `json:"usage"`
		} `json:"containers"`
	} `json:"items"`
}

var (
	// MetricsPodsCPU - CPU Gauge
	MetricsPodsCPU = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "kube_metrics_server_pods_cpu",
			Help: "Metrics Server Pods CPU",
		},
		[]string{"pod", "container"},
	)
	// MetricsPodsMEM - Memory Gauge
	MetricsPodsMEM = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "kube_metrics_server_pods_mem",
			Help: "Metrics Server Pods Memory",
		},
		[]string{"pod", "container"},
	)
)

// Collect responsible for get CPU and Memory data
func Collect() {

	var pods Info
	log.Println("Starting collect POD data,")

	re := regexp.MustCompile("[^0-9]")

	apiPod := api.Connect("pod")

	_ = json.NewDecoder(apiPod.Body).Decode(&pods)

	for i := range pods.Items {

		podName := pods.Items[i].Metadata.Name

		for j := range pods.Items[i].Containers {

			// Only numbers/String to float
			pods.Items[i].Containers[j].Usage.CPU = re.ReplaceAllLiteralString(pods.Items[i].Containers[j].Usage.CPU, "")
			CPUfloat, _ := strconv.ParseFloat(pods.Items[i].Containers[j].Usage.CPU, 64)

			// Only numbers/String to float
			pods.Items[i].Containers[j].Usage.Memory = re.ReplaceAllLiteralString(pods.Items[i].Containers[j].Usage.Memory, "")
			MEMfloat, _ := strconv.ParseFloat(pods.Items[i].Containers[j].Usage.Memory, 64)

			MetricsPodsCPU.With(prometheus.Labels{"pod": podName, "container": pods.Items[i].Containers[j].Name}).Add(CPUfloat)
			MetricsPodsMEM.With(prometheus.Labels{"pod": podName, "container": pods.Items[i].Containers[j].Name}).Add(MEMfloat)
		}

	}
	log.Println("POD data collected.")
}
