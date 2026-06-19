package sharedview

import (
	"fmt"
	sharedviewconsts "main/internal/delivery/view/shared/consts"
	"strings"
)

func (v *sharedView) NewUsefulKeys() string {
	var menus strings.Builder
	menus.WriteString("Useful keys:\n")
	for option := range int(sharedviewconsts.UsefulKeyQuit) + 1 {
		key := sharedviewconsts.UsefulKey(option)
		charToInput := sharedviewconsts.UsefulKeyCharInput[key]
		keyName := sharedviewconsts.UsefulKeyName[key]
		fmt.Fprintf(&menus, "  [%s] %s", charToInput, keyName)
		if key != sharedviewconsts.UsefulKeyQuit {
			menus.WriteString("\n")
		}
	}
	return menus.String()
}
