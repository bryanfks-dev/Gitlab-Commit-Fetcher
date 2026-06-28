package projectlistview

type MenuOption int

const (
	MenuOptionAddNewProject MenuOption = iota
	MenuOptionFetchProjectCommitMessages
	MenuOptionRemoveProject
	MenuOptionQuit
)

var MenuOptionCharInput = map[MenuOption]string{
	MenuOptionAddNewProject:              "1",
	MenuOptionFetchProjectCommitMessages: "2",
	MenuOptionRemoveProject:              "3",
	MenuOptionQuit:                       "q",
}

func (m MenuOption) MenuOptionCharInputs() []string {
	inputs := make([]string, 0)
	for option := range MenuOptionCharInput {
		inputs = append(inputs, MenuOptionCharInput[option])
	}
	return inputs
}

var MenuOptionName = map[MenuOption]string{
	MenuOptionAddNewProject:              "Add a new project",
	MenuOptionFetchProjectCommitMessages: "Fetch commit messages for a project",
	MenuOptionRemoveProject:              "Remove a project",
	MenuOptionQuit:                       "Quit",
}
