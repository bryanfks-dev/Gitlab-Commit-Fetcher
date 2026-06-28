package validatingtokenview

import (
	"context"
	"main/internal/delivery"
	usecase "main/internal/use_case"
)

type validatingTokenView struct {
	ctx context.Context
	uc  usecase.UseCase
}

type ValidatingTokenView interface {
	Render() (delivery.ViewState, error)
	Reset()
}

func NewValidatingTokenView(
	ctx context.Context,
	uc usecase.UseCase,
) ValidatingTokenView {
	return &validatingTokenView{
		ctx: ctx,
		uc:  uc,
	}
}
