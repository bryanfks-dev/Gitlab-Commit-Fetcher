package removeprojectview

import (
	"context"
	"main/internal/delivery"
	sharedview "main/internal/delivery/view/shared"
	"main/internal/service"
)

type removeProjectView struct {
	ctx    context.Context
	svc    service.Service
	shared sharedview.SharedView

	err       error
	projectID int
}

type RemoveProjectView interface {
	Render() (delivery.ViewState, error)
	Reset()
}

func NewRemoveProjectView(
	ctx context.Context,
	svc service.Service,
	shared sharedview.SharedView,
) RemoveProjectView {
	return &removeProjectView{
		ctx:    ctx,
		svc:    svc,
		shared: shared,
	}
}
