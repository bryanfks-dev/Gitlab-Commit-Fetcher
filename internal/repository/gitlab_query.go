package repository

import (
	"context"
	"fmt"
	"main/core/errors"
	httpclient "main/infra/http_client"
	httpclienterrors "main/infra/http_client/errors"
	"net/http"
)

func (r *repository) ValidateGitlabPersonalToken(
	ctx context.Context,
	token string,
) (bool, error) {
	r.logger.Infof("Validating Gitlab personal token")

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
		r.logger.Errorf("Validating Gitlab Personal Token failed: %v", err)
		return false, httpclienterrors.ErrSendRequestFailed
	}
	defer res.Body.Close()

	if res.StatusCode == http.StatusOK {
		r.logger.Infof("Gitlab Personal Token is valid")
		return true, nil
	}

	if res.StatusCode == http.StatusUnauthorized {
		r.logger.Errorf("Gitlab Personal Token is invalid: %d", res.StatusCode)
		return false, errors.ErrInvalidGitlabPersonalToken
	}

	r.logger.Warnf(
		"Unexpected status code when validating Gitlab Personal Token: %d",
		res.StatusCode,
	)
	return false, errors.ErrValidateGitlabPersonalTokenFailed
}

func (r *repository) GetGitlabProject() (string, error) {
	return "", nil
}
