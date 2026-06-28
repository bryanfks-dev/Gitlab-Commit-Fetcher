package addprojectview

import (
	"bufio"
	"fmt"
	"main/core/consts"
	"main/internal/delivery"
	deliveryutils "main/internal/delivery/utils"
	sharedviewconsts "main/internal/delivery/view/shared/consts"
	"os"
	"strconv"
	"time"

	"github.com/briandowns/spinner"
	"github.com/mitchellh/go-wordwrap"
)

func (v *addProjectView) Render() (delivery.ViewState, error) {
	var input string
	registeredProjectTable := v.shared.NewRegisteredProjectTable()
	getProjectIDMsg := wordwrap.WrapString(
		"Get project ID by visiting a project web URL, click three dots on the right side, then click `Copy project ID: xxx`.",
		consts.WindowMaxWidth,
	)
	usefulKeys := v.shared.NewUsefulKeys()

	for {
		deliveryutils.ClearScreen()

		fmt.Println(registeredProjectTable)
		fmt.Println()

		if v.projectID != 0 {
			nextState, err := v.renderRegisterGitLabProjectSpinner()
			if err != nil {
				v.err = err
				continue
			}
			return nextState, nil
		}

		fmt.Println(usefulKeys)
		fmt.Println()
		fmt.Println("[GUIDE]")
		fmt.Println(getProjectIDMsg)
		fmt.Println()

		if v.err != nil {
			fmt.Printf("* Error: %s\n", v.err.Error())
			fmt.Println()
		}

		fmt.Println("Enter project ID to add:")
		fmt.Print("> ")

		scanner := bufio.NewScanner(os.Stdin)
		if scanner.Scan() {
			input = scanner.Text()
			switch input {
			case sharedviewconsts.UsefulKeyCharInput[sharedviewconsts.UsefulKeyBack]:
				v.Reset()
				return delivery.ViewStateProjectList, nil
			case sharedviewconsts.UsefulKeyCharInput[sharedviewconsts.UsefulKeyQuit]:
				v.Reset()
				return delivery.ViewStateExit, nil
			}

			if err := v.ValidateProjectID(input); err != nil {
				v.err = err
				continue
			}

			// Project ID is always number at this point, so we can ignore the
			// error when converting to int
			v.projectID, _ = strconv.Atoi(input)
			continue
		}
	}
}

func (v *addProjectView) renderRegisterGitLabProjectSpinner() (delivery.ViewState, error) {
	spin := spinner.New(spinner.CharSets[9], 100*time.Millisecond)
	spin.Suffix = " Registering GitLab project..."
	spin.Start()
	err := v.uc.RegisterAndCheckProjectExistAndValid(
		v.ctx,
		v.projectID,
	)
	spin.Stop()

	v.Reset()
	if err != nil {
		return delivery.ViewStateAddProject, err
	}
	return delivery.ViewStateProjectList, nil
}
