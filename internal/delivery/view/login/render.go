package loginview

import (
	"fmt"
	"main/internal/delivery"
	deliveryutils "main/internal/delivery/utils"

	"github.com/rmasci/gopass"
)

func (v *loginView) Render() (delivery.ViewState, error) {
	generateTokenGuideMsg := v.newGenerateGitLabPersonalAccessTokenGuideMessage()

	for {
		deliveryutils.ClearScreen()

		fmt.Println("[GUIDE]")
		fmt.Println(generateTokenGuideMsg)
		fmt.Println()

		if v.err != nil {
			fmt.Printf("* Error: %s\n", v.err.Error())
			fmt.Println()
		}

		fmt.Println("Enter your GitLab personal access token to continue:")
		fmt.Print("> ")
		tokenBytes, err := gopass.GetPasswdMasked()
		if err != nil {
			v.Reset()
			return delivery.ViewStateLogin, err
		}

		v.token = string(tokenBytes)
		if err := v.ValidateGitLabPersonalAccessToken(v.token); err != nil {
			v.err = err
			continue
		}

		err = v.svc.SaveGitlabPersonalAccessToken(v.ctx, v.token)
		if err != nil {
			v.err = err
			continue
		}

		break
	}

	v.Reset()
	return delivery.ViewStateProjectList, nil
}
