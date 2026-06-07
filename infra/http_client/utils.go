package httpclient

import (
	"context"
	httpclientconsts "main/infra/http_client/consts"
	"time"
)

func newContextWithTimeoutIfNotSet(
	ctx context.Context,
) (context.Context, context.CancelFunc) {
	var cancel context.CancelFunc = func() {}

	if _, ok := ctx.Deadline(); !ok {
		ctx, cancel = context.WithTimeout(
			ctx,
			httpclientconsts.HttpClientDefaultTimeoutInSeconds*time.Second,
		)
	}
	return ctx, cancel
}
