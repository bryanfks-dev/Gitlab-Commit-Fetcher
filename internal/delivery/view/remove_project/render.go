package removeprojectview

import (
	"bufio"
	"fmt"
	"main/internal/delivery"
	deliveryutils "main/internal/delivery/utils"
	sharedviewconsts "main/internal/delivery/view/shared/consts"
	"os"
	"strconv"
	"time"

	"github.com/briandowns/spinner"
)

func (v *removeProjectView) Render() (delivery.ViewState, error) {
	var input string
	registeredProjectTable := v.shared.NewRegisteredProjectTable()
	usefulKeys := v.shared.NewUsefulKeys()

	for {
		deliveryutils.ClearScreen()

		fmt.Println(registeredProjectTable)
		fmt.Println()

		if v.projectID != 0 {
			nextState, err := v.renderRemoveGitLabProjectSpinner()
			if err != nil {
				v.err = err
				continue
			}
			return nextState, nil
		}

		fmt.Println(usefulKeys)
		fmt.Println()

		if v.err != nil {
			fmt.Printf("* Error: %s\n", v.err.Error())
			fmt.Println()
		}

		fmt.Println("Enter project ID to remove:")
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

func (v *removeProjectView) renderRemoveGitLabProjectSpinner() (delivery.ViewState, error) {
	spin := spinner.New(spinner.CharSets[9], 100*time.Millisecond)
	spin.Suffix = " Removing GitLab project..."
	spin.Start()
	err := v.svc.RemoveProject(v.ctx, v.projectID)
	spin.Stop()

	v.Reset()
	if err != nil {
		return delivery.ViewStateRemoveProject, err
	}
	return delivery.ViewStateProjectList, nil
}
