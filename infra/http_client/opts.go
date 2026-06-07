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

type HttpClientGetOptions = HttpClientOptions
type HttpClientPostOptions struct {
	HttpClientOptions
	Body []byte
}
