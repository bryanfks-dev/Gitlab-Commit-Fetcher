package validatingtokenview

import (
	"context"
	"main/internal/delivery"
	"main/internal/service"
)

type validatingTokenView struct {
	ctx context.Context
	svc service.Service
}

type ValidatingTokenView interface {
	Render() (delivery.ViewState, error)
	Reset()
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
