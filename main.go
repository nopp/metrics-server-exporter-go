package main

import (
	"fmt"
	"io/ioutil"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

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

	dat, err := ioutil.ReadFile("/var/run/secrets/kubernetes.io/serviceaccount/token")
	check(err)
	fmt.Print(string(dat))
}
