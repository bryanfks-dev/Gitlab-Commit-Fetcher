package loginview

import (
	"context"
	"main/internal/delivery"
	"main/internal/service"
)

type loginView struct {
	ctx           context.Context
	svc           service.Service
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
	svc service.Service,
	gitlabBaseURL string,
) LoginView {
	return &loginView{
		ctx:           ctx,
		svc:           svc,
		gitlabBaseURL: gitlabBaseURL,
	}
}
