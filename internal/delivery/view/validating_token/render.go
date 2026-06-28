package validatingtokenview

import (
	"errors"
	coreerrors "main/core/errors"
	"main/internal/delivery"
	"time"

	"github.com/briandowns/spinner"
)

func (v *validatingTokenView) Render() (delivery.ViewState, error) {
	spin := spinner.New(spinner.CharSets[11], 100*time.Millisecond)
	spin.Suffix = " Validating existing GitLab personal access token..."
	spin.Start()
	err := v.uc.CheckGitlabPersonalAccessTokenExistAndValid(v.ctx)
	spin.Stop()

	if err != nil {
		if errors.Is(err, coreerrors.ErrInvalidGitLabPersonalAccessToken) ||
			errors.Is(err, coreerrors.ErrGitLabPersonalAccessTokenNotFound) {
			return delivery.ViewStateLogin, nil
		}
		return delivery.ViewStateValidatingToken, err
	}
	return delivery.ViewStateProjectList, nil
}
