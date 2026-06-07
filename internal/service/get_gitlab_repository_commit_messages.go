package service

import (
	"context"
	"main/data/dto"
	vaultconsts "main/infra/vault/consts"
)

func (s *service) GetGitlabRepositoryCommitMessages(
	ctx context.Context,
	projectID int,
	lowerRef string,
	upperRef string,
) ([]string, error) {
	s.logger.Infof(
		"Getting Gitlab repository commit messages from %s to %s",
		lowerRef,
		upperRef,
	)

	token, err := s.vaultRepo.Get(vaultconsts.GitlabPersonalTokenKey)
	if err != nil {
		return []string{}, err
	}
	tokenStr := string(token)

	query := dto.GetGitlabRepositoryRefComparationQuery{
		Ctx:       ctx,
		ProjectID: projectID,
		Token:     tokenStr,
		LowerRef:  lowerRef,
		UpperRef:  upperRef,
	}
	res, err := s.repo.GetGitlabRefComparation(&query)
	if err != nil {
		return []string{}, err
	}

	commitMsgs := []string{}
	for _, commit := range res.Commits {
		commitMsgs = append(commitMsgs, commit.GetMessages()...)
	}
	return commitMsgs, nil
}
