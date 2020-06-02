package common

import (
	"crypto/tls"
	"crypto/x509"
	"io/ioutil"
	"net/http"
)

const (
	apiUrl = "https://kubernetes.default.svc/apis/metrics.k8s.io/v1beta1/namespaces/kube-system/pods"
	// apiURL = "https://kubernetes.default.svc/apis/metrics.k8s.io/v1beta1/nodes"
	token  = "/var/run/secrets/kubernetes.io/serviceaccount/token"
	caCert = "/var/run/secrets/kubernetes.io/serviceaccount/ca.crt"
)

func ConnectToAPI() *http.Response {

	var bearer = "Bearer " + string(returnDataFile(token))
	req, err := http.NewRequest("GET", apiUrl, nil)
	req.Header.Add("Authorization", bearer)

	caCertPool := x509.NewCertPool()
	caCertPool.AppendCertsFromPEM(returnDataFile(caCert))

	client := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				RootCAs: caCertPool,
			},
		},
	}

	resp, err := client.Do(req)
	if err != nil {
		panic("Can't connect to api.")
	}
	return resp
}

func returnDataFile(filePath string) []byte {

	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		panic(err)
	}
	return data
}
