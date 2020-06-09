package api

import (
	"crypto/tls"
	"crypto/x509"
	"io/ioutil"
	"log"
	"metrics-server-exporter-go/config"
	"net/http"
	"regexp"
	"strconv"
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
		panic(err)
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

// ReturnFloat - Remove non number caracter and parse to float64
func ReturnFloat(s string) float64 {

	re := regexp.MustCompile("[^0-9]")

	s = re.ReplaceAllLiteralString(s, "")
	numFloat, err := strconv.ParseFloat(s, 64)
	if err != nil {
		log.Fatal(err)
	}

	return numFloat
}
