package projectcommitmessagesview

import (
	"context"
	"main/internal/delivery"
	usecase "main/internal/use_case"
)

type projectCommitMessagesView struct {
	ctx   context.Context
	uc    usecase.UseCase
	state projectCommitMessagesState

	err      error
	lowerRef string
	upperRef string
}

type ProjectCommitMessagesView interface {
	Render() (delivery.ViewState, error)
	Reset()
}

func NewProjectCommitMessagesView(
	ctx context.Context,
	uc usecase.UseCase,
) ProjectCommitMessagesView {
	return &projectCommitMessagesView{
		ctx:   ctx,
		uc:    uc,
		state: projectCommitMessagesStateInputRefComprations,
	}
}
