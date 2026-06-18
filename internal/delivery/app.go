package delivery

import (
	"context"
	"errors"
	"main/core/config"
	coreerrors "main/core/errors"
	deliveryutils "main/internal/delivery/utils"
	addprojectview "main/internal/delivery/view/add_project"
	loginview "main/internal/delivery/view/login"
	projectlistview "main/internal/delivery/view/project_list"
	sharedview "main/internal/delivery/view/shared"
	validatingtokenview "main/internal/delivery/view/validating_token"
	"main/internal/service"
)

type app struct {
	ctx context.Context
	svc service.Service
	cfg *config.Config

	state           ViewState
	validatingToken validatingtokenview.ValidatingTokenView
	login           loginview.LoginView
	projectList     projectlistview.ProjectListView
	addProject      addprojectview.AddProjectView
}

type App interface {
	Run() error
}

func NewApp(ctx context.Context, svc service.Service, cfg *config.Config) App {
	shared := sharedview.NewSharedView(ctx, svc)

	return &app{
		ctx:             ctx,
		svc:             svc,
		cfg:             cfg,
		validatingToken: validatingtokenview.NewValidatingTokenView(ctx, svc),
		login:           loginview.NewLoginView(ctx, svc, cfg.Gitlab.BaseURL),
		projectList:     projectlistview.NewProjectListView(ctx, svc, shared),
		addProject:      addprojectview.NewAddProjectView(ctx, svc, shared),
	}
}

func (a *app) Run() error {
	a.state = ViewStateValidatingToken
	err := a.RenderView()

	if err != nil &&
		!errors.Is(err, coreerrors.ErrGitLabPersonalAccessTokenNotFound) &&
		!errors.Is(err, coreerrors.ErrInvalidGitLabPersonalAccessToken) {
		return err
	}

	if err != nil {
		a.state = ViewStateLogin
	}

	a.state = ViewStateProjectList
	return a.RenderView()
}

func (a *app) RenderView() error {
	deliveryutils.ClearScreen()

	switch a.state {
	case ViewStateLogin:
		return a.login.Render()
	case ViewStateValidatingToken:
		return a.validatingToken.Render()
	case ViewStateProjectList:
		return a.projectList.Render()
	case ViewStateAddProject:
		return a.addProject.Render()
	default:
		panic("unknown view state")
	}
}

func (a *app) UpdateViewState(state ViewState) {
	a.state = state
	a.RenderView()
}
