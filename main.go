package main

import (
	"fmt"
	"io/ioutil"
)

func returnString(filePath string) string {
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		panic(err)
	}
	return string(data)
}

const (
	api_url = "https://kubernetes.default.svc"
)

func main() {
	// url := "https://kubernetes.default.svc/"

	// // Create a Bearer string by appending string access token
	// var bearer = "Bearer " + <ACCESS TOKEN HERE>

	// // Create a new request using http
	// req, err := http.NewRequest("GET", url, nil)

	// // add authorization header to the req
	// req.Header.Add("Authorization", bearer)

	// // Send req using http Client
	// client := &http.Client{}
	// resp, err := client.Do(req)
	// if err != nil {
	//     log.Println("Error on response.\n[ERRO] -", err)
	// }

	// body, _ := ioutil.ReadAll(resp.Body)
	// log.Println(string([]byte(body)))

	token := returnString("/var/run/secrets/kubernetes.io/serviceaccount/token")
	ca := returnString("/var/run/secrets/kubernetes.io/serviceaccount/ca.crt")
	fmt.Print(token, ca, api_url)
}
