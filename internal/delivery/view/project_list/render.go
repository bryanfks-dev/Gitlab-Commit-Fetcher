package projectlistview

import (
	"bufio"
	"fmt"
	"main/internal/delivery"
	deliveryutils "main/internal/delivery/utils"
	"os"
	"strings"
)

func (v *projectListView) Render() (delivery.ViewState, error) {
	registeredProjectTable := v.shared.NewRegisteredProjectTable()
	menus := v.newMenuOptions()

	for {
		deliveryutils.ClearScreen()

		fmt.Println(registeredProjectTable)
		fmt.Println()
		fmt.Println(menus)
		fmt.Println()

		if v.err != nil {
			fmt.Printf("* Error: %s\n", v.err.Error())
			fmt.Println()
		}

		fmt.Print("> ")
		scanner := bufio.NewScanner(os.Stdin)
		if scanner.Scan() {
			v.selectedOption = scanner.Text()
			if err := v.ValidateMenuOption(v.selectedOption); err != nil {
				v.err = err
				continue
			}
			break
		}
	}

	return v.handleMenuInput(v.selectedOption)
}

func (v *projectListView) newMenuOptions() string {
	var menus strings.Builder
	menus.WriteString("What would you like to do?\n")
	for option := range int(MenuOptionQuit) + 1 {
		menuOption := MenuOption(option)
		charToInput := MenuOptionCharInput[menuOption]
		menuName := MenuOptionName[menuOption]
		fmt.Fprintf(&menus, "  [%s] %s", charToInput, menuName)
		if option != int(MenuOptionQuit) {
			menus.WriteString("\n")
		}
	}
	return menus.String()
}
