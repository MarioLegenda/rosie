package httpClient

import (
	"bytes"
	"crypto/tls"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

const ClientErrorType = 1
const NetworkErrorType = 2
const BootError = 3

type Header struct {
	Name  string
	Value string
}

type JsonRequest struct {
	Headers map[string]string
	Url     string
	Method  string
	Body    []byte
}

type RawResponse struct {
	Status int
	Body   []byte
}

type HttpClient struct {
	Client *http.Client
}

type IClientError interface {
	GetCode() int
	GetMessage() string
	GetRequest() *JsonRequest
	Type() string
}

// recoverable
type ClientError struct {
	Code    int
	Request *JsonRequest
	Message string
}

type NetworkError struct {
	Code    int
	Request *JsonRequest
	Message string
}

func (e *ClientError) GetCode() int {
	return e.Code
}

func (e *ClientError) GetMessage() string {
	return fmt.Sprintf("Code: %d, Message: %s", e.GetCode(), e.Message)
}

func (e *ClientError) GetRequest() *JsonRequest {
	return e.Request
}

func (e *ClientError) Type() string {
	return "client_error"
}

func (e *NetworkError) GetCode() int {
	return e.Code
}

func (e *NetworkError) GetMessage() string {
	return fmt.Sprintf("Code: %d, Message: %s", e.GetCode(), e.Message)
}

func (e *NetworkError) GetRequest() *JsonRequest {
	return e.Request
}

func (e *NetworkError) Type() string {
	return "network_error"
}

func NewHttpClient(config *tls.Config, idleConn int, handshakeTimeout time.Duration) (*HttpClient, error) {
	tr := &http.Transport{
		TLSClientConfig:     config,
		MaxConnsPerHost:     idleConn,
		TLSHandshakeTimeout: handshakeTimeout,
	}

	client := &http.Client{
		Timeout:   time.Second * time.Duration(120),
		Transport: tr,
	}

	return &HttpClient{Client: client}, nil
}

func (ac *HttpClient) MakeJsonRequest(r *JsonRequest) (RawResponse, IClientError) {
	request, err := http.NewRequest(r.Method, r.Url, bytes.NewBuffer(r.Body))

	request.SetBasicAuth("mgmotor", "g2062U8QWcER")

	if err != nil {
		return RawResponse{}, &NetworkError{
			Code:    BootError,
			Request: r,
			Message: "Could not create client",
		}
	}

	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Accepts", "application/json")

	if len(r.Headers) != 0 {
		for k, v := range r.Headers {
			request.Header.Set(k, v)
		}
	}

	response, err := ac.Client.Do(request)

	if err != nil {
		// An appErrors is returned if caused by client policy (such as CheckRedirect), or failure to speak HTTP
		// (such as a network connectivity problem). A non-2xx status code doesn't cause an appErrors.
		return RawResponse{}, &NetworkError{
			Code:    NetworkErrorType,
			Request: r,
			Message: fmt.Sprintf("Could not make client request with message: %s", err.Error()),
		}
	}

	defer response.Body.Close()

	// if an appErrors is nil, response.Body will always be at least []byte
	body, err := ioutil.ReadAll(response.Body)

	if err != nil {
		return RawResponse{}, &ClientError{
			Code:    ClientErrorType,
			Request: r,
			Message: "Unpacking JSON body failed",
		}
	}

	return RawResponse{
		Status: response.StatusCode,
		Body:   body,
	}, nil
}
