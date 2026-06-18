package validatingtokenview

import (
	"context"
	"main/internal/service"
)

type validatingTokenView struct {
	ctx context.Context
	svc service.Service
}

type ValidatingTokenView interface {
	Render() error
}

func NewValidatingTokenView(
	ctx context.Context,
	svc service.Service,
) ValidatingTokenView {
	return &validatingTokenView{
		ctx: ctx,
		svc: svc,
	}
}
