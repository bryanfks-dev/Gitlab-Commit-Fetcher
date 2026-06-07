package httpclient

import (
	"context"
	httpclientconsts "main/infra/http_client/consts"
	"time"
)

func newContextWithTimeoutIfNotSet(
	ctx context.Context,
	timeoutInSeconds *int,
) (context.Context, context.CancelFunc) {
	var cancel context.CancelFunc = func() {}
	if timeoutInSeconds == nil {
		t := httpclientconsts.HttpClientDefaultTimeoutInSeconds
		timeoutInSeconds = &t
	}

	if _, ok := ctx.Deadline(); !ok {
		ctx, cancel = context.WithTimeout(
			ctx,
			time.Duration(*timeoutInSeconds)*time.Second,
		)
	}
	return ctx, cancel
}
