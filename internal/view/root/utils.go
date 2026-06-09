package rootview

import (
	"context"
	"main/internal/service"
)

func determineInitialRootState(
	ctx context.Context,
	svc service.Service,
) rootState {
	ok, err := svc.CheckGitlabPersonalAccessTokenExistAndValid(ctx)
	if !ok || err != nil {
		return rootStateLogin
	}
	return rootStateMain
}
