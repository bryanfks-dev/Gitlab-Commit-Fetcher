package main

import (
	"context"
	"main/di"
	"main/internal/delivery"
	addprojectview "main/internal/delivery/view/add_project"
	loginview "main/internal/delivery/view/login"
	projectlistview "main/internal/delivery/view/project_list"
	sharedview "main/internal/delivery/view/shared"
	validatingtokenview "main/internal/delivery/view/validating_token"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	cont := di.NewContainer(ctx)
	app := cont.App

	sharedView := sharedview.NewSharedView(ctx, cont.Svc)
	validatingTokenView := validatingtokenview.NewValidatingTokenView(
		ctx,
		cont.Svc,
	)
	loginView := loginview.NewLoginView(ctx, cont.Svc, cont.Cfg.Gitlab.BaseURL)
	projectListView := projectlistview.NewProjectListView(
		ctx,
		cont.Svc,
		sharedView,
	)
	addProjectView := addprojectview.NewAddProjectView(
		ctx,
		cont.Svc,
		sharedView,
	)

	app.RegisterView(delivery.ViewStateValidatingToken, validatingTokenView)
	app.RegisterView(delivery.ViewStateLogin, loginView)
	app.RegisterView(delivery.ViewStateProjectList, projectListView)
	app.RegisterView(delivery.ViewStateAddProject, addProjectView)
	app.Run()
}
