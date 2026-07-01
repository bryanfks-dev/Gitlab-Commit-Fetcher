package projectcommitmessagesview

import (
	"context"
	"main/data/viewmodel"
	"main/internal/delivery"
	sharedview "main/internal/delivery/view/shared"
	"main/internal/service"
	usecase "main/internal/use_case"
)

type projectCommitMessagesView struct {
	ctx    context.Context
	svc    service.Service
	uc     usecase.UseCase
	state  projectCommitMessagesState
	shared sharedview.SharedView

	err                   error
	project               *viewmodel.ProjectViewModel
	lowerRef              string
	upperRef              string
	commitMessages        []string
	commitMessageFilename string
}

type ProjectCommitMessagesView interface {
	Render() (delivery.ViewState, error)
	Reset()
}

func NewProjectCommitMessagesView(
	ctx context.Context,
	svc service.Service,
	uc usecase.UseCase,
	shared sharedview.SharedView,
) ProjectCommitMessagesView {
	return &projectCommitMessagesView{
		ctx:    ctx,
		svc:    svc,
		uc:     uc,
		state:  projectCommitMessagesStateInputProjectID,
		shared: shared,
	}
}
