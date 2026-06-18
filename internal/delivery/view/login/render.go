package loginview

import (
	"fmt"
	deliveryutils "main/internal/delivery/utils"

	"github.com/rmasci/gopass"
)

func (v *loginView) Render() error {
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
			return err
		}

		token := string(tokenBytes)
		if err := v.ValidateGitLabPersonalAccessToken(token); err != nil {
			v.err = err
			continue
		}

		err = v.svc.SaveGitlabPersonalAccessToken(v.ctx, token)
		if err != nil {
			v.err = err
			continue
		}

		break
	}

	return nil
}
