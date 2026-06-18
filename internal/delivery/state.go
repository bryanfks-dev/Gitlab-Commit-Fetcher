package delivery

type ViewState int

const (
	ViewStateLogin ViewState = iota
	ViewStateProjectList
	ViewStateValidatingToken
	ViewStateAddProject
	ViewStateAddingProject
	ViewStateInputRefComparation
	ViewStateCommitList
)
