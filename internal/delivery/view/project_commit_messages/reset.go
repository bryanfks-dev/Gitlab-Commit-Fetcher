package projectcommitmessagesview

func (v *projectCommitMessagesView) updateStateReset() {
	v.err = nil
}

func (v *projectCommitMessagesView) Reset() {
	v.err = nil
	v.lowerRef = ""
	v.upperRef = ""
	v.project = nil
	v.commitMessages = nil
	v.state = projectCommitMessagesStateInputProjectID
}
