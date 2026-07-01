package service

import (
	"context"

	"golang.design/x/clipboard"
)

func (s *service) CopyTextToClipboard(ctx context.Context, text string) {
	s.logger.Infof("Copying text to clipboard")

	content := []byte(text)
	clipboard.Write(clipboard.FmtText, content)
}
