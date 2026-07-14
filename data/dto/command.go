package dto

type WriteCommitMessageToFileCommand struct {
	ProjectID int
	LowerRef  string
	UpperRef  string
	Filename  string
	Messages  []string
	Prefix    *string
	Suffix    *string
}
