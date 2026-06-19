package sharedviewconsts

type UsefulKey int

const (
	UsefulKeyBack UsefulKey = iota
	UsefulKeyQuit
)

var UsefulKeyCharInput = map[UsefulKey]string{
	UsefulKeyBack: "b",
	UsefulKeyQuit: "q",
}

func (u UsefulKey) UsefulKeyCharInputs() []string {
	inputs := make([]string, 0)
	for option := range UsefulKeyCharInput {
		inputs = append(inputs, UsefulKeyCharInput[option])
	}
	return inputs
}

var UsefulKeyName = map[UsefulKey]string{
	UsefulKeyBack: "Back",
	UsefulKeyQuit: "Quit",
}
