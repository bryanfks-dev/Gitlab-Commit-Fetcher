package mainview

import (
	"context"
	"main/internal/service"

	"charm.land/bubbles/v2/textinput"
	tea "charm.land/bubbletea/v2"
)

type MainViewModel struct {
	ctx context.Context
	svc service.Service

	gitlabProjectErr        error
	gitlabProjectRefDiffErr error
	gitlabProject           string
	gitlabProjectRefDiff    textinput.Model
}

func NewMainViewModel(
	ctx context.Context,
	svc service.Service,
) MainViewModel {
	refDiff := textinput.New()
	refDiff.Placeholder = "x.x.x - x.x.x"
	refDiff.Prompt = "Old Ref - New Ref: "
	refDiff.Focus()
	refDiff.CharLimit = 255

	return MainViewModel{
		ctx: ctx,
		svc: svc,
	}
}

func (vm MainViewModel) Init() tea.Cmd {
	return nil
}

func (vm MainViewModel) Update(msg tea.Msg) (MainViewModel, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyPressMsg:
		switch msg.String() {
		case "enter":
		}
	}

	var cmd tea.Cmd
	vm.gitlabProjectRefDiff, cmd = vm.gitlabProjectRefDiff.Update(msg)
	return vm, cmd
}

func (vm MainViewModel) View() tea.View {
	return tea.NewView("")
}
