package textinputcomponent

import (
	"main/data/props"

	"charm.land/lipgloss/v2"
)

func NewTextInput(props props.TextInputProps) string {
	nonNilComp := make([]string, 0)
	if props.Label != nil {
		label := newTextInputLabel(*props.Label)
		nonNilComp = append(nonNilComp, label)
	}

	input := lipgloss.NewStyle().
		Border(lipgloss.RoundedBorder()).
		Padding(0, 2).
		SetString(props.Model.View()).
		Render()
	nonNilComp = append(nonNilComp, input)

	if props.Err != nil {
		err := newTextInputError(props.Err)
		nonNilComp = append(nonNilComp, err)
	}

	return lipgloss.JoinVertical(lipgloss.Left, nonNilComp...)
}
