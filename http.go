package main

import (
	"crypto/tls"
	"log"
	"simulation/httpClient"
)

func sendRequest(http http, url string) bool {
	response, clientError := http.client.MakeJsonRequest(&httpClient.JsonRequest{
		Url:    url,
		Method: "GET",
	})

	if clientError != nil {
		return false
	}

	if response.Status < 200 && response.Status > 299 {
		return false
	}

	return true
}

type http struct {
	client *httpClient.HttpClient
}

func newHttp() http {
	c, err := httpClient.NewHttpClient(&tls.Config{InsecureSkipVerify: true}, 1024, 0)

	if err != nil {
		log.Fatalln(err)
	}

	return http{
		client: c,
	}
}
