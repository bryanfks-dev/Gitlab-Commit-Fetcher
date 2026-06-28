package projectcommitmessagesview

import "main/internal/delivery"

func (v *projectCommitMessagesView) Render() (delivery.ViewState, error) {
	for {
		switch v.state {
		case projectCommitMessagesStateInputRefComprations:
			return v.renderInputRefComparations()
		case projectCommitMessagesStateDisplayCommitMessages:
			return v.renderDisplayCommitMessages()
		}
		panic("Invalid state")
	}
}

func (v *projectCommitMessagesView) renderInputRefComparations() (delivery.ViewState, error) {
}

func (v *projectCommitMessagesView) renderDisplayCommitMessages() (delivery.ViewState, error) {}
