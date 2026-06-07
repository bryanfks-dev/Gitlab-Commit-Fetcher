package handler

import (
	"bufio"
	"context"
	"fmt"
	"os"
	"strings"
)

func (h *handler) askAndSaveForGitlabPersonalAccessTokenUntilValid(
	ctx context.Context,
) {
	scanner := bufio.NewScanner(os.Stdin)
	var token string

	for {
		fmt.Print("Enter your Gitlab Personal Access Token: ")
		if !scanner.Scan() {
			h.logger.Errorf("Error reading token input: %v", scanner.Err())
			fmt.Println("Error reading token, try again")
			continue
		}

		token = strings.TrimSpace(scanner.Text())
		if err := h.validateGitlabPersonalToken(token); err != nil {
			fmt.Println(err.Error())
			continue
		}
		if err := h.svc.SaveGitlabPersonalAccessToken(ctx, token); err != nil {
			h.logger.Printf("Error saving token: %v\n", err)
			continue
		}
		break
	}
}
