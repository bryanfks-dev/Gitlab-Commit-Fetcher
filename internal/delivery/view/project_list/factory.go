package projectlistview

import (
	"context"
	"main/internal/delivery"
	sharedview "main/internal/delivery/view/shared"
	"main/internal/service"
)

type projectListView struct {
	ctx    context.Context
	svc    service.Service
	shared sharedview.SharedView
	err    error

	selectedOption string
}

type ProjectListView interface {
	Render() (delivery.ViewState, error)
	Reset()
}

func NewProjectListView(
	ctx context.Context,
	svc service.Service,
	sharedView sharedview.SharedView,
) ProjectListView {
	return &projectListView{
		ctx:    ctx,
		svc:    svc,
		shared: sharedView,
	}
}
