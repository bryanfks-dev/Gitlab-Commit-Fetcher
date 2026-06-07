package httpclient

import (
	"bytes"
	"io"
	"net/http"
)

func (h *httpClient) Get(opts *HttpClientGetOptions) (*http.Response, error) {
	ctx, cancel := newContextWithTimeoutIfNotSet(opts.Ctx)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, opts.URL, nil)
	if err != nil {
		return nil, err
	}
	req.Header = opts.Header

	return h.client.Do(req)
}

func (h *httpClient) Post(opts *HttpClientPostOptions) (*http.Response, error) {
	ctx, cancel := newContextWithTimeoutIfNotSet(opts.Ctx)
	defer cancel()

	bodyBytes := io.NopCloser(bytes.NewReader(opts.Body))
	req, err := http.NewRequestWithContext(
		ctx,
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
