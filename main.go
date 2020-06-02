package main

import (
	"crypto/tls"
	"crypto/x509"
	"io/ioutil"
	"log"
	"net/http"
)

func returnDataFile(filePath string) []byte {
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		panic(err)
	}
	return data
}

const (
	apiURL = "https://kubernetes.default.svc/apis/metrics.k8s.io/v1beta1/pods"
	token  = "/var/run/secrets/kubernetes.io/serviceaccount/token"
	caCert = "/var/run/secrets/kubernetes.io/serviceaccount/ca.crt"
)

func main() {

	// Create a Bearer string by appending string access token
	var bearer = "Bearer " + string(returnDataFile(token))

	// Create a new request using http
	req, err := http.NewRequest("GET", apiURL, nil)

	// add authorization header to the req
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
		log.Println("Error on response.\n[ERRO] -", err)
	}

	body, _ := ioutil.ReadAll(resp.Body)
	log.Println(string([]byte(body)))
}
