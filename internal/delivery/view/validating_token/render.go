package validatingtokenview

import (
	"time"

	"github.com/briandowns/spinner"
)

func (v *validatingTokenView) Render() error {
	spin := spinner.New(spinner.CharSets[11], 100*time.Millisecond)
	spin.Suffix = " Validating existing GitLab personal access token..."
	spin.Start()
	err := v.svc.CheckGitlabPersonalAccessTokenExistAndValid(v.ctx)
	spin.Stop()
	return err
}
