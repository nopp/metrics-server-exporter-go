package pod

import (
	"encoding/json"
	"log"
	"metrics-server-exporter-go/api"
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
		[]string{"pod_name", "pod_namespace", "pod_container_name"},
	)
	// MetricsPodsMEM - Memory Gauge
	MetricsPodsMEM = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "kube_metrics_server_pods_mem",
			Help: "Metrics Server Pods Memory",
		},
		[]string{"pod_name", "pod_namespace", "pod_container_name"},
	)
)

// Collect responsible for get CPU and Memory data
func Collect() {

	var pods Info
	log.Println("Starting collect POD data,")

	apiPod := api.Connect("pod")

	_ = json.NewDecoder(apiPod.Body).Decode(&pods)

	for i := range pods.Items {

		podName := pods.Items[i].Metadata.Name
		podNamespace := pods.Items[i].Metadata.Namespace

		for j := range pods.Items[i].Containers {

			MetricsPodsCPU.With(prometheus.Labels{"pod_name": podName, "pod_namespace": podNamespace, "pod_container_name": pods.Items[i].Containers[j].Name}).Add(api.ReturnFloat(pods.Items[i].Containers[j].Usage.CPU))
			MetricsPodsMEM.With(prometheus.Labels{"pod_name": podName, "pod_namespace": podNamespace, "pod_container_name": pods.Items[i].Containers[j].Name}).Add(api.ReturnFloat(pods.Items[i].Containers[j].Usage.Memory))
		}

	}
	log.Println("POD data collected.")
}
