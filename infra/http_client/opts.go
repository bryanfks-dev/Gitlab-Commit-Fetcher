package httpclient

import (
	"context"
	"net/http"
)

type HttpClientOptions struct {
	Ctx    context.Context
	URL    string
	Header http.Header
}

type HttpClientWithBodyOptions struct {
	HttpClientOptions
	Body []byte
}

type HttpClientGetOptions = HttpClientOptions
type HttpClientPostOptions = HttpClientWithBodyOptions
type HttpClientPutOptions = HttpClientWithBodyOptions
type HttpClientDeleteOptions = HttpClientOptions
