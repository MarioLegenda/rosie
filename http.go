package main

import (
	"crypto/tls"
	"log"
	"net/http"
	"simulation/httpClient"
)

type mainHttpClient struct {
	client *http.Client
}

func newHttp() mainHttpClient {
	client := httpClient.NewClient(httpClient.ClientParams{
		Transport: &http.Transport{
			TLSClientConfig:     &tls.Config{InsecureSkipVerify: true},
			MaxConnsPerHost:     1024,
			TLSHandshakeTimeout: 0,
		},
		CheckRedirect: nil,
		Jar:           nil,
		Timeout:       0,
	})

	return mainHttpClient{client: client}
}

func sendRequest(http mainHttpClient, url string) (*http.Response, error) {
	request, err := httpClient.NewRequest(httpClient.Request{
		Headers: nil,
		Url:     url,
		Method:  "GET",
		Body:    nil,
	})

	if err != nil {
		log.Fatalln(err)
	}

	response, err := httpClient.Make(request, http.client)

	if err != nil {
		log.Fatalln(err)
	}

	return response, nil
}
