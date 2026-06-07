package httpclient

import (
	"net/http"
)

type httpClient struct {
	client *http.Client
}

type HttpClient interface {
	Get(opts *HttpClientGetOptions) (*http.Response, error)
	Post(opts *HttpClientPostOptions) (*http.Response, error)
}

func NewHttpClient() HttpClient {
	client := http.Client{}

	return &httpClient{
		client: &client,
	}
}
