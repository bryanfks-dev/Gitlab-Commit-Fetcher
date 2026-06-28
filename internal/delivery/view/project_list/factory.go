package projectlistview

import (
	"context"
	"main/internal/delivery"
	sharedview "main/internal/delivery/view/shared"
	usecase "main/internal/use_case"
)

type projectListView struct {
	ctx    context.Context
	uc     usecase.UseCase
	shared sharedview.SharedView

	err            error
	selectedOption string
}

type ProjectListView interface {
	Render() (delivery.ViewState, error)
	Reset()
}

func NewProjectListView(
	ctx context.Context,
	uc usecase.UseCase,
	sharedView sharedview.SharedView,
) ProjectListView {
	return &projectListView{
		ctx:    ctx,
		uc:     uc,
		shared: sharedView,
	}
}
