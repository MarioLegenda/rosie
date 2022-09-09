package main

import (
	"crypto/tls"
	"log"
	"therebelsource/simulation/httpClient"
)

func sendRequest(url string) bool {
	client, err := httpClient.NewHttpClient(&tls.Config{InsecureSkipVerify: true})

	if err != nil {
		log.Fatalln(err)
	}

	response, clientError := client.MakeJsonRequest(&httpClient.JsonRequest{
		Url:    url,
		Method: "GET",
	})

	if clientError != nil {
		log.Fatalln(clientError)
	}

	if response.Status != 200 {
		return false
	}

	return true
}
