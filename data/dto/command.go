package dto

type WriteCommitMessageToFileCommand struct {
	LowerRef string
	UpperRef string
	Filename string
	Messages []string
	Prefix   *string
	Suffix   *string
}
