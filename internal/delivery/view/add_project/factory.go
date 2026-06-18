package addprojectview

import (
	"context"
	sharedview "main/internal/delivery/view/shared"
	"main/internal/service"
)

type addProjectView struct {
	ctx    context.Context
	svc    service.Service
	shared sharedview.SharedView
	err    error
}

type AddProjectView interface {
	Render() error
}

func NewAddProjectView(ctx context.Context, svc service.Service, shared sharedview.SharedView) AddProjectView {
	return &addProjectView{
		ctx:    ctx,
		svc:    svc,
		shared: shared,
	}
}
