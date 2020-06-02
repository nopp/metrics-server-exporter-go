package common

import "io/ioutil"

const (
	APIUrl = "https://kubernetes.default.svc/apis/metrics.k8s.io/v1beta1/namespaces/kube-system/pods"
	// apiURL = "https://kubernetes.default.svc/apis/metrics.k8s.io/v1beta1/nodes"
	Token  = "/var/run/secrets/kubernetes.io/serviceaccount/token"
	CACert = "/var/run/secrets/kubernetes.io/serviceaccount/ca.crt"
)

func ReturnDataFile(filePath string) []byte {
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		panic(err)
	}
	return data
}
