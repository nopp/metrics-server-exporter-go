# metrics-server-exporter-go (Go Version)

![Go](https://github.com/nopp/metrics-server-exporter-go/workflows/Go/badge.svg)
![Docker](https://github.com/nopp/metrics-server-exporter-go/workflows/Docker/badge.svg)
[![Total alerts](https://img.shields.io/lgtm/alerts/g/nopp/metrics-server-exporter-go.svg?logo=lgtm&logoWidth=18)](https://lgtm.com/projects/g/nopp/metrics-server-exporter-go/alerts/)
[![Language grade: Go](https://img.shields.io/lgtm/grade/go/g/nopp/metrics-server-exporter-go.svg?logo=lgtm&logoWidth=18)](https://lgtm.com/projects/g/nopp/metrics-server-exporter-go/context:go)

Based on https://github.com/grupozap/metrics-server-exporter (Python Version)

metrics-server-exporter-go provides cpu and memory metrics for nodes and pods, directly querying the metrics-server API `/apis/metrics.k8s.io/v1beta1/{pods, nodes}`

### Node metrics

* kube_metrics_server_nodes_mem
	* Provides nodes memory information in kibibytes.
* kube_metrics_server_nodes_cpu
	* Provides nodes CPU information in nanocores.

##### labels

* instance

### Pod metrics

* kube_metrics_server_pods_mem
	* Provides pods/container memory information.
* kube_metrics_server_pods_cpu
	* Provides pods/container memory information.

##### labels

* pod_name
* pod_container_name
