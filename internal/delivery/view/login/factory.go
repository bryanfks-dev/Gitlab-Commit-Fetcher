package loginview

import (
	"context"
	"main/internal/service"
)

type loginView struct {
	ctx           context.Context
	svc           service.Service
	gitlabBaseURL string
	err           error
}

type LoginView interface {
	Render() error
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
