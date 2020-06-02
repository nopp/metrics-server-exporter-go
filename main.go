package main

import (
	"encoding/json"
	"log"
	"net/http"

	"metrics-server-exporter-go/common"
	"metrics-server-exporter-go/pod"

	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {

	apiGo := common.ConnectToAPI()

	var pods pod.Info
	_ = json.NewDecoder(apiGo.Body).Decode(&pods)

	log.Println(pods)

	http.Handle("/metrics", promhttp.Handler())
	http.ListenAndServe(":8484", nil)
}
