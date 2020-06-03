package prometheus

import (
	"encoding/json"
	"log"
	"metrics-server-exporter-go/metrics"
	"metrics-server-exporter-go/node"
	"metrics-server-exporter-go/pod"
)

func Collect() {

	// POD
	apiPod := metrics.ConnectToAPI("pod")

	var pods pod.Info
	_ = json.NewDecoder(apiPod.Body).Decode(&pods)

	// for i := range pods.Items {

	// 	podName := pods.Items[i].Metadata.Name
	// 	podNamespace := pods.Items[i].Metadata.Namespace

	// 	fmt.Println(podName, podNamespace)
	// 	fmt.Println(pods.Items[i].Containers)
	// }
	log.Println(pods)

	// NODE
	apiNode := metrics.ConnectToAPI("node")

	var nodes node.Info
	_ = json.NewDecoder(apiNode.Body).Decode(&nodes)

	// for i := range nodes.Items {

	// 	nodeName := nodes.Items[i].Metadata.Name

	// 	fmt.Println(nodeName)
	// }

	log.Println(nodes)
}
