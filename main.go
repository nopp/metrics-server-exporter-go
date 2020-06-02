package main

import (
	"crypto/tls"
	"crypto/x509"
	"encoding/json"
	"log"
	"net/http"
	"time"

	"metrics-server-exporter-go/common"

	"github.com/prometheus/client_golang/prometheus/promhttp"
)

type Pod struct {
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

type Node struct {
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

// func returnDataFile(filePath string) []byte {
// 	data, err := ioutil.ReadFile(filePath)
// 	if err != nil {
// 		panic(err)
// 	}
// 	return data
// }

// const (
// 	apiURL = "https://kubernetes.default.svc/apis/metrics.k8s.io/v1beta1/namespaces/kube-system/pods"
// 	// apiURL = "https://kubernetes.default.svc/apis/metrics.k8s.io/v1beta1/nodes"
// 	token  = "/var/run/secrets/kubernetes.io/serviceaccount/token"
// 	caCert = "/var/run/secrets/kubernetes.io/serviceaccount/ca.crt"
// )

func main() {

	// Create a Bearer string by appending string access token
	var bearer = "Bearer " + string(common.ReturnDataFile(common.Token))

	// Create a new request using http
	req, err := http.NewRequest("GET", common.APIUrl, nil)

	// add authorization header to the req
	req.Header.Add("Authorization", bearer)

	caCertPool := x509.NewCertPool()
	caCertPool.AppendCertsFromPEM(common.ReturnDataFile(common.CACert))

	client := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				RootCAs: caCertPool,
			},
		},
	}

	resp, err := client.Do(req)
	if err != nil {
		log.Println("Can't connect to api.\n[ERRO] -", err)
	}

	// body, _ := ioutil.ReadAll(resp.Body)
	// log.Println(string([]byte(body)))

	// var nodes Node
	var pods Pod
	// _ = json.NewDecoder(resp.Body).Decode(&nodes)
	_ = json.NewDecoder(resp.Body).Decode(&pods)

	log.Println(pods.Items[0].Metadata.Name, pods.Items[0].Containers[0].Name, pods.Items[0].Containers[0].Usage.CPU, pods.Items[0].Containers[0].Usage.Memory)

	http.Handle("/metrics", promhttp.Handler())
	http.ListenAndServe(":8484", nil)
}
