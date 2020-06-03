package api

import (
	"crypto/tls"
	"crypto/x509"
	"io/ioutil"
	"metrics-server-exporter-go/config"
	"net/http"
)

// Connect - Responsible to connect to node/pod API
func Connect(option string) *http.Response {

	var apiURL string

	if option == "pod" {
		apiURL = config.APIURLPods
	} else {
		apiURL = config.APIURLNodes
	}

	var bearer = "Bearer " + string(returnDataFile(config.Token))
	req, err := http.NewRequest("GET", apiURL, nil)
	if err != nil {
		panic("Can't create new request.")
	}
	req.Header.Add("Authorization", bearer)

	caCertPool := x509.NewCertPool()
	caCertPool.AppendCertsFromPEM(returnDataFile(config.CACert))

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
