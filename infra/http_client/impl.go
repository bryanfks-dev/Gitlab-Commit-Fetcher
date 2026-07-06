package httpclient

import (
	"bytes"
	"io"
	"net/http"
)

func (h *httpClient) Get(opts *HttpClientGetOptions) (*http.Response, error) {
	req, err := http.NewRequestWithContext(
		opts.Ctx,
		http.MethodGet,
		opts.URL,
		nil,
	)
	if err != nil {
		return nil, err
	}
	req.Header = opts.Header

	return h.client.Do(req)
}

func (h *httpClient) Post(opts *HttpClientPostOptions) (*http.Response, error) {
	bodyBytes := io.NopCloser(bytes.NewReader(opts.Body))
	req, err := http.NewRequestWithContext(
		opts.Ctx,
		http.MethodPost,
		opts.URL,
		bodyBytes,
	)
	if err != nil {
		return nil, err
	}

	req.Header = opts.Header

	return h.client.Do(req)
}

func (h *httpClient) Put(opts *HttpClientPostOptions) (*http.Response, error) {
	bodyBytes := io.NopCloser(bytes.NewReader(opts.Body))
	req, err := http.NewRequestWithContext(
		opts.Ctx,
		http.MethodPut,
		opts.URL,
		bodyBytes,
	)
	if err != nil {
		return nil, err
	}

	req.Header = opts.Header

	return h.client.Do(req)
}

func (h *httpClient) Patch(
	opts *HttpClientPostOptions,
) (*http.Response, error) {
	bodyBytes := io.NopCloser(bytes.NewReader(opts.Body))
	req, err := http.NewRequestWithContext(
		opts.Ctx,
		http.MethodPatch,
		opts.URL,
		bodyBytes,
	)
	if err != nil {
		return nil, err
	}

	req.Header = opts.Header

	return h.client.Do(req)
}

func (h *httpClient) Delete(
	opts *HttpClientDeleteOptions,
) (*http.Response, error) {
	req, err := http.NewRequestWithContext(
		opts.Ctx,
		http.MethodDelete,
		opts.URL,
		nil,
	)
	if err != nil {
		return nil, err
	}
	req.Header = opts.Header

	return h.client.Do(req)
}
