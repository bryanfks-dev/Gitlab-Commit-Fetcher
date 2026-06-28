package projectlistview

import "main/internal/delivery"

func (v *projectListView) handleMenuInput(
	input string,
) (delivery.ViewState, error) {
	switch input {
	case MenuOptionCharInput[MenuOptionAddNewProject]:
		v.Reset()
		return delivery.ViewStateAddProject, nil
	case MenuOptionCharInput[MenuOptionFetchProjectCommitMessages]:
		v.Reset()
		return delivery.ViewStateProjectCommits, nil
	case MenuOptionCharInput[MenuOptionRemoveProject]:
		v.Reset()
		return delivery.ViewStateRemoveProject, nil
	case MenuOptionCharInput[MenuOptionQuit]:
		v.Reset()
		return delivery.ViewStateExit, nil
	}
	panic("Invalid input option")
}
