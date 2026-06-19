package delivery

type View interface {
	Render() (ViewState, error)
	Reset()
}
