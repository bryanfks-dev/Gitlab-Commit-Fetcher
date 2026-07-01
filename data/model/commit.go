package model

import (
	"fmt"
	"strings"
)

type CommitModel struct {
	ID              string
	Message         string
	Author          string
	CommitParentIDs []string
}

func NewCommitModel(
	id string,
	message string,
	author string,
	commitParentIDs []string,
) *CommitModel {
	return &CommitModel{
		ID:              id,
		Message:         message,
		Author:          author,
		CommitParentIDs: commitParentIDs,
	}
}

func (m *CommitModel) IsMergeCommit() bool {
	return len(m.CommitParentIDs) > 1
}

func (m *CommitModel) GetMessages() []string {
	msg := strings.TrimSpace(m.Message)
	if msg == "" {
		return []string{}
	}

	parentParts := strings.SplitN(msg, "\n\n", 2)
	if len(parentParts) == 1 {
		return []string{strings.TrimSpace(parentParts[0])}
	}

	childParts := strings.Split(parentParts[1], "\n")
	msgs := []string{}
	for _, part := range childParts {
		if strings.TrimSpace(part) == "" {
			continue
		}
		msg := fmt.Sprintf("[%s] %s", parentParts[0], strings.TrimSpace(part))
		msgs = append(msgs, msg)
	}
	return msgs
}
