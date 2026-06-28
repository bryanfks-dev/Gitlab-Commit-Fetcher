package delivery

type ViewState int

const (
	ViewStateLogin ViewState = iota
	ViewStateProjectList
	ViewStateValidatingToken
	ViewStateAddProject
	ViewStateRemoveProject
	ViewStateProjectCommits
	ViewStateExit
)
