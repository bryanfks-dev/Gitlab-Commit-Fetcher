package addprojectview

import (
	"bufio"
	"fmt"
	"main/core/consts"
	deliveryutils "main/internal/delivery/utils"
	"os"

	"github.com/mitchellh/go-wordwrap"
)

func (v *addProjectView) Render() error {
	var input string
	registeredProjectTable := v.shared.NewRegisteredProjectTable()
	getProjectIDMsg := wordwrap.WrapString(
		"Get project ID by visiting a project web URL, click three dots on the right side, then click `Copy project ID: xxx`.",
		consts.WindowMaxWidth,
	)

	for {
		deliveryutils.ClearScreen()

		fmt.Println(registeredProjectTable)
		fmt.Println()
		fmt.Println("[GUIDE]")
		fmt.Println(getProjectIDMsg)
		fmt.Println()

		if v.err != nil {
			fmt.Printf("* Error: %s\n", v.err.Error())
			fmt.Println()
		}

		fmt.Println("Enter project ID:")
		fmt.Print("> ")

		scanner := bufio.NewScanner(os.Stdin)
		if scanner.Scan() {
			input = scanner.Text()
			if err := v.ValidateProjectID(input); err != nil {
				v.err = err
				continue
			}
			break
		}
	}

	return nil
}
