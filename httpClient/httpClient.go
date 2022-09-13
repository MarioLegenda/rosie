package httpClient

import (
	"bytes"
	"net/http"
	"time"
)

type Header struct {
	Name  string
	Value string
}

type Request struct {
	Headers map[string]string
	Url     string
	Method  string
	Body    []byte
}

type ClientParams struct {
	Transport     http.RoundTripper
	CheckRedirect func(req *http.Request, via []*http.Request) error
	Jar           http.CookieJar
	Timeout       time.Duration
}

func NewClient(params ClientParams) *http.Client {
	return &http.Client{
		Transport:     params.Transport,
		CheckRedirect: params.CheckRedirect,
		Jar:           params.Jar,
		Timeout:       params.Timeout,
	}
}

func NewRequest(request Request) (*http.Request, error) {
	r, err := http.NewRequest(request.Method, request.Url, bytes.NewBuffer(request.Body))

	if err != nil {
		return nil, err
	}

	if len(request.Headers) != 0 {
		for k, v := range request.Headers {
			r.Header.Set(k, v)
		}
	}

	return r, nil
}

func Make(request *http.Request, client *http.Client) (*http.Response, error) {
	return client.Do(request)
}
