package loginview

import (
	"context"
	"main/internal/delivery"
	usecase "main/internal/use_case"
)

type loginView struct {
	ctx           context.Context
	uc            usecase.UseCase
	gitlabBaseURL string

	err   error
	token string
}

type LoginView interface {
	Render() (delivery.ViewState, error)
	Reset()
}

func NewLoginView(
	ctx context.Context,
	uc usecase.UseCase,
	gitlabBaseURL string,
) LoginView {
	return &loginView{
		ctx:           ctx,
		uc:            uc,
		gitlabBaseURL: gitlabBaseURL,
	}
}
