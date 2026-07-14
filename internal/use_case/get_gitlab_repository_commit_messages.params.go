package usecase

type WriteCommitMessagesToFileOptions struct {
	ProjectID  int
	LowerRef   string
	UpperRef   string
	Filename   string
	CommitMsgs []string
}
