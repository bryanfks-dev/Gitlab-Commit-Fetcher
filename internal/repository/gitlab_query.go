package repository

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

func (r *repository) ValidateGitlabPersonalAccessToken(
	ctx context.Context,
	token string,
) (bool, error) {
	r.logger.Infof("Validating Gitlab personal access token")

	url := fmt.Sprintf("%s/user", r.gitlabAPIBaseURL)
	header := http.Header{
		"PRIVATE-TOKEN": []string{token},
	}

	opts := httpclient.HttpClientGetOptions{
		Ctx:    ctx,
		URL:    url,
		Header: header,
	}
	res, err := r.httpClient.Get(&opts)
	if err != nil {
		r.logger.Errorf(
			"Validating Gitlab personal access token failed: %v",
			err,
		)
		return false, httpclienterrors.ErrSendRequestFailed
	}
	defer res.Body.Close()

	if res.StatusCode == http.StatusUnauthorized {
		return false, errors.ErrInvalidGitlabPersonalAccessToken
	}

	if res.StatusCode != http.StatusOK {
		r.logger.Warnf(
			"Unexpected status code when validating Gitlab personal access token: %d",
			res.StatusCode,
		)
		return false, errors.ErrValidateGitlabPersonalAccessTokenFailed
	}
	return true, nil
}

func (r *repository) GetGitlabRefComparation(
	query *dto.GetGitlabRepositoryRefComparationQuery,
) (*dto.RefComparationResultDTO, error) {
	r.logger.Infof(
		"Getting Gitlab repository with projectID `%d` tag comparation between `%s` and `%s`",
		query.ProjectID,
		query.LowerRef,
		query.UpperRef,
	)

	url := fmt.Sprintf(
		"%s/projects/%d/repository/compare?from=%s&to=%s",
		r.gitlabAPIBaseURL,
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
	res, err := r.httpClient.Get(&opts)
	if err != nil {
		r.logger.Errorf(
			"Getting Gitlab repository ref comparation failed: %v",
			err,
		)
		return nil, errors.ErrGetGitlabRepositoryRefComparationFailed
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		r.logger.Warnf(
			"Unexpected status code when getting Gitlab repository ref comparation: %d",
			res.StatusCode,
		)
		return nil, errors.ErrGetGitlabRepositoryRefComparationFailed
	}

	var decodedRes response.GitlabRefComparationResponse
	if err := json.NewDecoder(res.Body).Decode(&decodedRes); err != nil {
		r.logger.Errorf(
			"Decoding Gitlab repository ref comparation response failed: %v",
			err,
		)
		return nil, errors.ErrGetGitlabRepositoryRefComparationFailed
	}
	dto := mapper.MapGitlabRefComparationToRefComparationResultDTO(&decodedRes)
	return dto, nil
}

func (r *repository) GetGitlabProject(
	ctx context.Context,
	projectID int,
	token string,
) (*model.ProjectModel, error) {
	r.logger.Infof("Getting Gitlab project")

	url := fmt.Sprintf("%s/projects/%d", r.gitlabAPIBaseURL, projectID)
	header := http.Header{
		"PRIVATE-TOKEN": []string{token},
	}

	opts := httpclient.HttpClientGetOptions{
		Ctx:    ctx,
		URL:    url,
		Header: header,
	}
	res, err := r.httpClient.Get(&opts)
	if err != nil {
		r.logger.Errorf("Getting Gitlab project failed: %v", err)
		return nil, errors.ErrGetProjectFailed
	}
	defer res.Body.Close()

	if res.StatusCode == http.StatusNotFound {
		r.logger.Warnf("Gitlab project with ID `%d` not found", projectID)
		return nil, nil
	}

	if res.StatusCode != http.StatusOK {
		r.logger.Warnf(
			"Unexpected status code when getting Gitlab project: %d",
			res.StatusCode,
		)
		return nil, errors.ErrGetProjectFailed
	}

	var decodedRes response.GitlabProjectResponse
	if err := json.NewDecoder(res.Body).Decode(&decodedRes); err != nil {
		r.logger.Errorf("Decoding Gitlab project response failed: %v", err)
		return nil, errors.ErrGetProjectFailed
	}
	model := mapper.MapGitlabProjectToProjectModel(&decodedRes)
	return model, nil
}
