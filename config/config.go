package config

const (
	// K8s Config
	APIURLPods  = "https://kubernetes.default.svc/apis/metrics.k8s.io/v1beta1/pods"
	APIURLNodes = "https://kubernetes.default.svc/apis/metrics.k8s.io/v1beta1/nodes"
	Token       = "/var/run/secrets/kubernetes.io/serviceaccount/token"
	CACert      = "/var/run/secrets/kubernetes.io/serviceaccount/ca.crt"

	// APP Config
	Host = "0.0.0.0"
	Port = "8484"
)
