package projectlistview

import (
	"context"
	sharedview "main/internal/delivery/view/shared"
	"main/internal/service"
)

type projectListView struct {
	ctx    context.Context
	svc    service.Service
	shared sharedview.SharedView
	err    error
}

type ProjectListView interface {
	Render() error
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
