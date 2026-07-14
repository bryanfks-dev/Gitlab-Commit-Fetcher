package projectcommitmessagesview

import (
	"fmt"
	"main/data/mapper"
	"main/internal/delivery"
	"strings"
)

func (v *projectCommitMessagesView) handleMenuInput() (delivery.ViewState, error) {
	switch v.state {
	case projectCommitMessagesStateInputProjectID:
		return v.renderInputProjectID()
	case projectCommitMessagesStateInputRefComprations:
		return v.renderInputRefComparations()
	case projectCommitMessagesStateDisplayCommitMessages:
		return v.renderDisplayCommitMessages()
	}
	panic("Invalid state")
}

func (v *projectCommitMessagesView) fetchProject(projectID int) error {
	project, err := v.uc.GetRegisteredGitlabProject(projectID)
	if err != nil {
		return err
	}

	v.project = mapper.MapProjectModelToProjectViewModel(project)
	return nil
}

func (v *projectCommitMessagesView) fetchCommitMessages() error {
	filename, commitMessages, err := v.uc.GetGitlabRepositoryCommitMessages(
		v.ctx,
		v.project.ID,
		v.lowerRef,
		v.upperRef,
	)
	if err != nil {
		return err
	}

	v.commitMessages = commitMessages
	v.commitMessageFilename = filename
	return nil
}

func (v *projectCommitMessagesView) copyCommitMessagesToClipboard() {
	var commitMsgs strings.Builder
	for index, commit := range v.commitMessages {
		if index == len(v.commitMessages)-1 {
			fmt.Fprintf(&commitMsgs, "- %s", commit)
			continue
		}
		fmt.Fprintf(&commitMsgs, "- %s\n", commit)
	}
	v.svc.CopyTextToClipboard(v.ctx, commitMsgs.String())
}
