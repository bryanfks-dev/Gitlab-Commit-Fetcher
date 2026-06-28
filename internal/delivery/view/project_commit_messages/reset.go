package projectcommitmessagesview

func (v *projectCommitMessagesView) Reset() {
	v.err = nil
	v.lowerRef = ""
	v.upperRef = ""
}
