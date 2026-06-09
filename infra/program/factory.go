package program

import (
	tea "charm.land/bubbletea/v2"
)

type program struct {
	Prog *tea.Program
}

type Program interface {
	Start() error
}

func NewProgram(initModel tea.Model) Program {
	prog := tea.NewProgram(initModel)
	return &program{
		Prog: prog,
	}
}
