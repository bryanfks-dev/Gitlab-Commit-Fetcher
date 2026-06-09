package rootview

import (
	"context"
	"main/internal/service"
	loginview "main/internal/view/login"
	mainview "main/internal/view/main"

	tea "charm.land/bubbletea/v2"
)

type rootState int

const (
	rootStateLogin rootState = iota
	rootStateMain
	rootStateCommits
)

type RootViewModel struct {
	loginVM loginview.LoginViewModel
	mainVM  mainview.MainViewModel

	state rootState
}

func NewRootViewModel(
	ctx context.Context,
	svc service.Service,
	loginVM loginview.LoginViewModel,
	mainVM mainview.MainViewModel,
) RootViewModel {
	state := determineInitialRootState(ctx, svc)
	return RootViewModel{
		state:   state,
		loginVM: loginVM,
		mainVM:  mainVM,
	}
}

func (vm RootViewModel) Init() tea.Cmd {
	if vm.state == rootStateLogin {
		return vm.loginVM.Init()
	}
	return vm.mainVM.Init()
}

func (vm RootViewModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyPressMsg:
		switch msg.String() {
		case "ctrl+c":
			return vm, tea.Quit
		}
	}

	switch vm.state {
	case rootStateLogin:
		var cmd tea.Cmd
		vm.loginVM, cmd = vm.loginVM.Update(msg)
		return vm, cmd
	case rootStateMain:
		var cmd tea.Cmd
		vm.mainVM, cmd = vm.mainVM.Update(msg)
		return vm, cmd
	}

	return vm, nil
}

func (vm RootViewModel) View() tea.View {
	switch vm.state {
	case rootStateLogin:
		return vm.loginVM.View()
	case rootStateMain:
		return vm.mainVM.View()
	}
	return tea.NewView("")
}
