package projectlistview

import "main/internal/delivery"

func (v *projectListView) handleMenuInput(
	input string,
) (delivery.ViewState, error) {
	switch input {
	case MenuOptionCharInput[MenuOptionAddNewProject]:
		v.Reset()
		return delivery.ViewStateAddProject, nil
	case MenuOptionCharInput[MenuOptionFetchProjectCommits]:
		v.Reset()
		return delivery.ViewStateCommitList, nil
	case MenuOptionCharInput[MenuOptionRemoveProject]:
		v.Reset()
		return delivery.ViewStateDeleteProject, nil
	case MenuOptionCharInput[MenuOptionQuit]:
		return delivery.ViewStateExit, nil
	}
	panic("Invalid input option")
}
