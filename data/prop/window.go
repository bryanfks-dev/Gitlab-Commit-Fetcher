package prop

type WindowProps struct {
	Width  *int
	Height *int
}

func NewWindowProps(width *int, height *int) *WindowProps {
	return &WindowProps{
		Width:  width,
		Height: height,
	}
}
