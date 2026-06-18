package projectlistview

import (
	"bufio"
	"fmt"
	deliveryutils "main/internal/delivery/utils"
	"os"
	"strings"
)

func (v *projectListView) Render() error {
	registeredProjectTable := v.shared.NewRegisteredProjectTable()
	menus := v.newMenuOptions()
	var input string

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
			input = scanner.Text()
			if err := v.ValidateMenuOption(input); err != nil {
				v.err = err
				continue
			}
			break
		}
	}

	v.handleMenuInput(input)
	return nil
}

func (v *projectListView) handleMenuInput(input string) {
	switch input {
	case MenuOptionCharInput[MenuOptionAddNewProject]:
	case MenuOptionCharInput[MenuOptionFetchProjectCommits]:
	case MenuOptionCharInput[MenuOptionRemoveProject]:
	case MenuOptionCharInput[MenuOptionQuit]:
		os.Exit(0)
	}
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
