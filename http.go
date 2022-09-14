package main

import (
	"crypto/tls"
	"net/http"
	"net/http/httptrace"
	"simulation/httpClient"
	"time"
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

func sendRequest(url string) (*http.Response, time.Duration, error) {
	request, err := httpClient.NewRequest(httpClient.Request{
		Headers: nil,
		Url:     url,
		Method:  "GET",
		Body:    nil,
	})

	if err != nil {
		return nil, 0, err
	}

	var start time.Time
	var end time.Duration
	trace := &httptrace.ClientTrace{
		DNSStart: func(dsi httptrace.DNSStartInfo) {
			start = time.Now()
		},
		GotFirstResponseByte: func() {
			end = time.Since(start)
		},
	}

	request = request.WithContext(httptrace.WithClientTrace(request.Context(), trace))

	response, err := http.DefaultTransport.RoundTrip(request)

	if err != nil {
		return nil, 0, err
	}

	return response, end, nil
}
