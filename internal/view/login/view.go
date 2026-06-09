package loginview

import (
	"context"
	"main/core/consts"
	"main/data/props"
	textinputcomponent "main/internal/component/text_input"
	"main/internal/service"
	logincomponent "main/internal/view/login/component"

	"charm.land/bubbles/v2/textinput"
	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"
)

type LoginViewModel struct {
	ctx           context.Context
	gitlabBaseURL string
	svc           service.Service

	err                       error
	gitlabPersonalAccessToken textinput.Model
}

func NewLoginViewModel(
	ctx context.Context,
	gitlabBaseURL string,
	svc service.Service,
) LoginViewModel {
	tokenIn := textinput.New()
	tokenIn.Placeholder = "xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx"
	tokenIn.Prompt = "GitLab Personal Access Token: "
	tokenIn.Focus()
	tokenIn.SetWidth(60)
	tokenIn.CharLimit = 64

	return LoginViewModel{
		ctx:                       ctx,
		gitlabBaseURL:             gitlabBaseURL,
		svc:                       svc,
		gitlabPersonalAccessToken: tokenIn,
	}
}

func (vm LoginViewModel) Init() tea.Cmd {
	return textinput.Blink
}

func (vm LoginViewModel) Update(msg tea.Msg) (LoginViewModel, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyPressMsg:
		switch msg.String() {
		case "enter":
			if err := vm.handleOnInputSubmit(); err != nil {
				vm.err = err
				vm.gitlabPersonalAccessToken.SetValue("")
			}
		}
	}

	var cmd tea.Cmd
	vm.gitlabPersonalAccessToken, cmd = vm.gitlabPersonalAccessToken.Update(msg)
	return vm, cmd
}

func (vm LoginViewModel) View() tea.View {
	retrievalGuide := logincomponent.GetTokenRetrievalGuide(vm.gitlabBaseURL)
	info := logincomponent.GetInformation()
	inputGuide := logincomponent.GetInputGuide()

	inputProps := props.TextInputProps{
		Err:   vm.err,
		Model: vm.gitlabPersonalAccessToken,
	}
	input := textinputcomponent.NewTextInput(inputProps)

	s := lipgloss.JoinVertical(
		lipgloss.Left,
		retrievalGuide,
		input,
		inputGuide,
		info,
	)
	cvs := lipgloss.NewStyle().
		Margin(0, 2).
		Width(consts.WindowWidth).
		SetString(s).
		Render()
	return tea.NewView(cvs)
}
