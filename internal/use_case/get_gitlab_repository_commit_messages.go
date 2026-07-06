package usecase

import (
	"context"
	"fmt"
	datastructures "main/core/data_structures"
	"main/data/dto"
	vaultconsts "main/infra/vault/consts"
	"time"
)

type WriteCommitMessagesToFileOptions struct {
	LowerRef   string
	UpperRef   string
	Filename   string
	CommitMsgs []string
}

func (uc *useCase) GetGitlabRepositoryCommitMessages(
	ctx context.Context,
	projectID int,
	lowerRef string,
	upperRef string,
) (string, []string, error) {
	uc.logger.Infof(
		"Getting Gitlab repository commit messages from %s to %s",
		lowerRef,
		upperRef,
	)

	token, err := uc.vaultRepo.Get(vaultconsts.GitlabPersonalTokenKey)
	if err != nil {
		return "", []string{}, err
	}
	tokenStr := string(token)

	query := dto.GetGitlabRepositoryRefComparationQuery{
		Ctx:       ctx,
		ProjectID: projectID,
		Token:     tokenStr,
		LowerRef:  lowerRef,
		UpperRef:  upperRef,
	}
	res, err := uc.svc.GetGitlabRefComparation(&query)
	if err != nil {
		return "", []string{}, err
	}

	commitMsgs := datastructures.NewSet[string]()
	for _, commit := range res.Commits {
		if commit.IsMergeCommit() || commit.IsRevertCommit() {
			continue
		}
		for _, msg := range commit.GetMessages() {
			commitMsgs.Add(msg)
		}
	}

	filenameSuffix := uc.getCommitMessageFilenameSuffix()
	commitMsgSlice := commitMsgs.ToSlice()
	options := WriteCommitMessagesToFileOptions{
		Filename:   filenameSuffix,
		CommitMsgs: commitMsgSlice,
		LowerRef:   lowerRef,
		UpperRef:   upperRef,
	}
	err = uc.writeCommitMessagesToFile(options)
	if err != nil {
		return "", []string{}, err
	}

	return filenameSuffix, commitMsgSlice, nil
}

func (uc *useCase) writeCommitMessagesToFile(
	options WriteCommitMessagesToFileOptions,
) error {
	msgPrefix := "-"
	cmd := dto.WriteCommitMessageToFileCommand{
		Filename: options.Filename,
		Messages: options.CommitMsgs,
		LowerRef: options.LowerRef,
		UpperRef: options.UpperRef,
		Prefix:   &msgPrefix,
	}
	return uc.repo.WriteCommitMessageToFile(cmd)
}

func (uc *useCase) getCommitMessageFilenameSuffix() string {
	now := time.Now()
	unix := now.Unix()
	return fmt.Sprintf("%d", unix)
}
