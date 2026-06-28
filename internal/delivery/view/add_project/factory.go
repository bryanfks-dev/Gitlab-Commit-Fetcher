package addprojectview

import (
	"context"
	"main/internal/delivery"
	sharedview "main/internal/delivery/view/shared"
	usecase "main/internal/use_case"
)

type addProjectView struct {
	ctx    context.Context
	uc     usecase.UseCase
	shared sharedview.SharedView

	err       error
	projectID int
}

type AddProjectView interface {
	Render() (delivery.ViewState, error)
	Reset()
}

func NewAddProjectView(
	ctx context.Context,
	uc usecase.UseCase,
	shared sharedview.SharedView,
) AddProjectView {
	return &addProjectView{
		ctx:    ctx,
		uc:     uc,
		shared: shared,
	}
}
