package service

import (
	"context"
	"encoding/json"
	"fmt"
	"main/core/errors"
	"main/data/dto"
	"main/data/mapper"
	"main/data/model"
	"main/data/response"
	httpclient "main/infra/http_client"
	httpclienterrors "main/infra/http_client/errors"
	"net/http"
)

func (s *service) ValidateGitlabPersonalAccessToken(
	ctx context.Context,
	token string,
) error {
	s.logger.Infof("Validating Gitlab personal access token")

	url := fmt.Sprintf("%s/user", s.gitlabAPIBaseURL)
	header := http.Header{
		"PRIVATE-TOKEN": []string{token},
	}

	opts := httpclient.HttpClientGetOptions{
		Ctx:    ctx,
		URL:    url,
		Header: header,
	}
	res, err := s.httpClient.Get(&opts)
	if err != nil {
		s.logger.Errorf(
			"Validating Gitlab personal access token failed: %v",
			err,
		)
		return httpclienterrors.ErrSendRequestFailed
	}
	defer res.Body.Close()

	if res.StatusCode == http.StatusUnauthorized {
		return errors.ErrInvalidGitLabPersonalAccessToken
	}

	if res.StatusCode != http.StatusOK {
		s.logger.Warnf(
			"Unexpected status code when validating Gitlab personal access token: %d",
			res.StatusCode,
		)
		return errors.ErrValidateGitLabPersonalAccessTokenFailed
	}
	return nil
}

func (s *service) GetGitlabRefComparation(
	query *dto.GetGitlabRepositoryRefComparationQuery,
) (*dto.RefComparationResultDTO, error) {
	s.logger.Infof(
		"Getting Gitlab repository with projectID `%d` tag comparation between `%s` and `%s`",
		query.ProjectID,
		query.LowerRef,
		query.UpperRef,
	)

	url := fmt.Sprintf(
		"%s/projects/%d/repository/compare?from=%s&to=%s",
		s.gitlabAPIBaseURL,
		query.ProjectID,
		query.LowerRef,
		query.UpperRef,
	)
	header := http.Header{
		"PRIVATE-TOKEN": []string{query.Token},
	}

	opts := httpclient.HttpClientGetOptions{
		Ctx:    query.Ctx,
		URL:    url,
		Header: header,
	}
	res, err := s.httpClient.Get(&opts)
	if err != nil {
		s.logger.Errorf(
			"Getting Gitlab repository ref comparation failed: %v",
			err,
		)
		return nil, errors.ErrGetGitLabRepositoryRefComparationFailed
	}
	defer res.Body.Close()

	if res.StatusCode == http.StatusNotFound {
		s.logger.Warnf(
			"Gitlab repository with projectID `%d` not found",
			query.ProjectID,
		)
		return nil, errors.ErrInvalidGitLabRefComparation
	}
	if res.StatusCode != http.StatusOK {
		s.logger.Warnf(
			"Unexpected status code when getting Gitlab repository ref comparation: %d",
			res.StatusCode,
		)
		return nil, errors.ErrGetGitLabRepositoryRefComparationFailed
	}

	var decodedRes response.GitlabRefComparationResponse
	if err := json.NewDecoder(res.Body).Decode(&decodedRes); err != nil {
		s.logger.Errorf(
			"Decoding Gitlab repository ref comparation response failed: %v",
			err,
		)
		return nil, errors.ErrGetGitLabRepositoryRefComparationFailed
	}
	dto := mapper.MapGitlabRefComparationToRefComparationResultDTO(&decodedRes)
	for _, msg := range dto.Commits {
		s.logger.Infof(msg.Message)
	}
	return dto, nil
}

func (s *service) GetGitlabProject(
	ctx context.Context,
	projectID int,
	token string,
) (*model.ProjectModel, error) {
	s.logger.Infof("Getting Gitlab project")

	url := fmt.Sprintf("%s/projects/%d", s.gitlabAPIBaseURL, projectID)
	header := http.Header{
		"PRIVATE-TOKEN": []string{token},
	}

	opts := httpclient.HttpClientGetOptions{
		Ctx:    ctx,
		URL:    url,
		Header: header,
	}
	res, err := s.httpClient.Get(&opts)
	if err != nil {
		s.logger.Errorf("Getting Gitlab project failed: %v", err)
		return nil, errors.ErrGetProjectFailed
	}
	defer res.Body.Close()

	if res.StatusCode == http.StatusNotFound {
		s.logger.Warnf("Gitlab project with ID `%d` not found", projectID)
		return nil, nil
	}

	if res.StatusCode != http.StatusOK {
		s.logger.Warnf(
			"Unexpected status code when getting Gitlab project: %d",
			res.StatusCode,
		)
		return nil, errors.ErrGetProjectFailed
	}

	var decodedRes response.GitlabProjectResponse
	if err := json.NewDecoder(res.Body).Decode(&decodedRes); err != nil {
		s.logger.Errorf("Decoding Gitlab project response failed: %v", err)
		return nil, errors.ErrGetProjectFailed
	}
	model := mapper.MapGitlabProjectToProjectModel(&decodedRes)
	return model, nil
}
