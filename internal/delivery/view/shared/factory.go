package sharedview

import (
	"context"
	"main/internal/service"
)

type SharedView struct {
	ctx context.Context
	svc service.Service
}

func NewSharedView(ctx context.Context, svc service.Service) SharedView {
	return SharedView{
		ctx: ctx,
		svc: svc,
	}
}
