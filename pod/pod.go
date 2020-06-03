package pod

import (
	"encoding/json"
	"fmt"
	"log"
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
	MetricsPodsCPU = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "kube_metrics_server_pods_cpu",
			Help: "Metrics Server Pods CPU",
		},
		[]string{"pod", "container"},
	)
	MetricsPodsMEM = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "kube_metrics_server_pods_mem",
			Help: "Metrics Server Pods Memory",
		},
		[]string{"pod", "container"},
	)
)

func Collect() {

	re := regexp.MustCompile("[^0-9]")

	apiPod := api.Connect("pod")

	var pods Info
	_ = json.NewDecoder(apiPod.Body).Decode(&pods)

	for i := range pods.Items {

		podName := pods.Items[i].Metadata.Name
		podNamespace := pods.Items[i].Metadata.Namespace

		fmt.Println(podName, podNamespace)
		fmt.Println(pods.Items[i].Containers)

		for j := range pods.Items[i].Containers {
			pods.Items[i].Containers[j].Usage.CPU = re.ReplaceAllLiteralString(pods.Items[i].Containers[j].Usage.CPU, "")
			CPUfloat, _ := strconv.ParseFloat(pods.Items[i].Containers[j].Usage.CPU, 64)
			log.Println(podName, pods.Items[i].Containers[j].Name, CPUfloat)
			pods.Items[i].Containers[j].Usage.Memory = re.ReplaceAllLiteralString(pods.Items[i].Containers[j].Usage.Memory, "")
			MEMfloat, _ := strconv.ParseFloat(pods.Items[i].Containers[j].Usage.Memory, 64)
			MetricsPodsCPU.WithLabelValues(podName, pods.Items[i].Containers[j].Name).Add(CPUfloat)
			MetricsPodsMEM.WithLabelValues(podName, pods.Items[i].Containers[j].Name).Add(MEMfloat)
		}

	}
}
