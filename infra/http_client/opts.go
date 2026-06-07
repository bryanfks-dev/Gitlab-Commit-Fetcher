package httpclient

import (
	"context"
	"net/http"
)

type httpClientOptions struct {
	Ctx              context.Context
	URL              string
	Header           http.Header
	TimeoutInSeconds *int
}

type httpClientWithBodyOptions struct {
	httpClientOptions
	Body []byte
}

type HttpClientGetOptions = httpClientOptions
type HttpClientPostOptions = httpClientWithBodyOptions
type HttpClientPutOptions = httpClientWithBodyOptions
type HttpClientDeleteOptions = httpClientOptions
