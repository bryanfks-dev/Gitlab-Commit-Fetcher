package sharedview

import (
	"context"
	usecase "main/internal/use_case"
)

type sharedView struct {
	ctx context.Context
	uc  usecase.UseCase
}

type SharedView interface {
	NewRegisteredProjectTable() string
	NewUsefulKeys() string
}

func NewSharedView(ctx context.Context, uc usecase.UseCase) SharedView {
	return &sharedView{
		ctx: ctx,
		uc:  uc,
	}
}
