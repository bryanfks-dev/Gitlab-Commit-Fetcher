package loginview

import (
	"fmt"
	"strings"
)

func (vm LoginViewModel) handleOnInputSubmit() error {
	value := strings.Trim(vm.gitlabPersonalAccessToken.Value(), " ")
	if value == "" {
		return fmt.Errorf("Token cannot be empty")
	}
	return vm.svc.SaveGitlabPersonalAccessToken(vm.ctx, value)
}
