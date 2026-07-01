package repository

import (
	"fmt"
	"main/core/errors"
	"main/core/utils"
	"main/data/dto"
	"os"
	"strings"
)

func (r *repository) WriteCommitMessageToFile(
	cmd dto.WriteCommitMessageToFileCommand,
) error {
	filepath, err := utils.NewCommitMessagesFilepath(cmd.Filename)
	if err != nil {
		r.logger.Errorf("failed to create commit messages filepath: %v", err)
		return errors.ErrWriteCommitMessageFailed
	}

	parts := strings.Split(filepath, "/")
	if len(parts) > 1 {
		dir := strings.Join(parts[:len(parts)-1], "/")
		if err := os.MkdirAll(dir, 0755); err != nil {
			r.logger.Errorf(
				"failed to create directory for commit messages file: %v",
				err,
			)
			return errors.ErrCreateCommitMessageFileFailed
		}
	}

	title := fmt.Sprintf("[%s-%s]", cmd.LowerRef, cmd.UpperRef)
	msgs := []string{title}
	for _, msg := range cmd.Messages {
		if cmd.Prefix != nil {
			msg = fmt.Sprintf("%s %s", *cmd.Prefix, msg)
		}
		if cmd.Suffix != nil {
			msg = fmt.Sprintf("%s %s", msg, *cmd.Suffix)
		}
		msgs = append(msgs, msg)
	}

	msgWithNewLine := strings.Join(msgs, "\n")
	r.logger.Infof(msgWithNewLine)
	content := []byte(msgWithNewLine)
	if err := os.WriteFile(filepath, content, 0644); err != nil {
		r.logger.Errorf("failed to write commit messages to file: %v", err)
		return errors.ErrWriteCommitMessageFailed
	}
	return nil
}
