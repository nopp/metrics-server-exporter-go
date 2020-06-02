package main

import (
	"encoding/json"
	"log"
	"net/http"

	"metrics-server-exporter-go/common"
	"metrics-server-exporter-go/node"
	"metrics-server-exporter-go/pod"

	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {

	// // Create a Bearer string by appending string access token
	// var bearer = "Bearer " + string(common.ReturnDataFile(common.Token))

	// // Create a new request using http
	// req, err := http.NewRequest("GET", common.APIUrl, nil)

	// // add authorization header to the req
	// req.Header.Add("Authorization", bearer)

	// caCertPool := x509.NewCertPool()
	// caCertPool.AppendCertsFromPEM(common.ReturnDataFile(common.CACert))

	// client := &http.Client{
	// 	Transport: &http.Transport{
	// 		TLSClientConfig: &tls.Config{
	// 			RootCAs: caCertPool,
	// 		},
	// 	},
	// }

	// resp, err := client.Do(req)
	// if err != nil {
	// 	log.Println("Can't connect to api.\n[ERRO] -", err)
	// }

	apiGo := common.ConnectToAPI()

	// body, _ := ioutil.ReadAll(resp.Body)
	// log.Println(string([]byte(body)))

	var nodes node.Node
	var pods pod.Pod
	// _ = json.NewDecoder(resp.Body).Decode(&nodes)
	// _ = json.NewDecoder(resp.Body).Decode(&pods)
	_ = json.NewDecoder(apiGo.Body).Decode(&nodes)
	_ = json.NewDecoder(apiGo.Body).Decode(&pods)

	log.Println(pods.Items[0].Metadata.Name, pods.Items[0].Containers[0].Name, pods.Items[0].Containers[0].Usage.CPU, pods.Items[0].Containers[0].Usage.Memory)
	log.Println(nodes.Items[0].Metadata.Name, nodes.Items[0].Usage.CPU, nodes.Items[0].Usage.Memory)

	http.Handle("/metrics", promhttp.Handler())
	http.ListenAndServe(":8484", nil)
}
