package removeprojectview

import (
	"context"
	"main/internal/delivery"
	sharedview "main/internal/delivery/view/shared"
	usecase "main/internal/use_case"
)

type removeProjectView struct {
	ctx    context.Context
	uc     usecase.UseCase
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
	uc usecase.UseCase,
	shared sharedview.SharedView,
) RemoveProjectView {
	return &removeProjectView{
		ctx:    ctx,
		uc:     uc,
		shared: shared,
	}
}
