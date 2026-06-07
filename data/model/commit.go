package model

import (
	"fmt"
	"strings"
)

type CommitModel struct {
	ID      string
	Message string
	Author  string
}

func (m *CommitModel) GetMessages() []string {
	parentParts := strings.SplitN(m.Message, "\n\n", 2)
	if len(parentParts) == 1 {
		return []string{strings.TrimSpace(parentParts[0])}
	}

	childParts := strings.Split(parentParts[1], "\n")
	msgs := []string{}
	for _, part := range childParts {
		msg := fmt.Sprintf(
			"[%s] %s",
			parentParts[0],
			strings.TrimSpace(part),
		)
		msgs = append(msgs, msg)
	}
	return msgs
}
