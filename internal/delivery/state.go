package delivery

type ViewState int

const (
	ViewStateLogin ViewState = iota
	ViewStateProjectList
	ViewStateValidatingToken
	ViewStateAddProject
	ViewStateDeleteProject
	ViewStateInputRefComparation
	ViewStateCommitList
	ViewStateExit
)
