package sharedview

import (
	"context"
	"main/internal/service"
)

type sharedView struct {
	ctx context.Context
	svc service.Service
}

type SharedView interface {
	NewRegisteredProjectTable() string
	NewUsefulKeys() string
}

func NewSharedView(ctx context.Context, svc service.Service) SharedView {
	return &sharedView{ctx: ctx,
		svc: svc,
	}
}
