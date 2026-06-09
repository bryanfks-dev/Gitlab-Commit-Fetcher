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
	ctx     context.Context
	svc     service.Service
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
	return RootViewModel{
		state:   rootStateLogin,
		svc:     svc,
		loginVM: loginVM,
		mainVM:  mainVM,
	}
}

func (vm RootViewModel) Init() tea.Cmd {
	ok, err := vm.svc.CheckGitlabPersonalAccessTokenExistAndValid(vm.ctx)
	if err != nil {
		panic(err)
	}

	if !ok {
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
