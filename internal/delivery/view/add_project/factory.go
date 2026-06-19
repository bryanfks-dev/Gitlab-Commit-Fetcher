package addprojectview

import (
	"context"
	"main/internal/delivery"
	sharedview "main/internal/delivery/view/shared"
	"main/internal/service"
)

type addProjectView struct {
	ctx    context.Context
	svc    service.Service
	shared sharedview.SharedView
	err    error

	projectID int
}

type AddProjectView interface {
	Render() (delivery.ViewState, error)
	Reset()
}

func NewAddProjectView(
	ctx context.Context,
	svc service.Service,
	shared sharedview.SharedView,
) AddProjectView {
	return &addProjectView{
		ctx:    ctx,
		svc:    svc,
		shared: shared,
	}
}
