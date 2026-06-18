package projectlistview

import (
	"fmt"
	"slices"
)

func (v *projectListView) ValidateMenuOption(option string) error {
	charInputs := MenuOptionQuit.MenuOptionCharInputs()
	if !slices.Contains(charInputs, option) {
		return fmt.Errorf("Invalid menu option: %s", option)
	}
	return nil
}
