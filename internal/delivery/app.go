package delivery

import (
	"bufio"
	"context"
	"fmt"
	"main/core/config"
	deliveryutils "main/internal/delivery/utils"
	usecase "main/internal/use_case"
	"os"
)

type app struct {
	ctx context.Context
	uc  usecase.UseCase
	cfg *config.Config

	views        map[ViewState]View
	currentState ViewState
	history      []ViewState
}

type App interface {
	RegisterView(state ViewState, view View)
	Run()
}

func NewApp(ctx context.Context, uc usecase.UseCase, cfg *config.Config) App {
	return &app{
		ctx:          ctx,
		uc:           uc,
		cfg:          cfg,
		currentState: ViewStateValidatingToken,
	}
}

func (a *app) RegisterView(state ViewState, view View) {
	if a.views == nil {
		a.views = make(map[ViewState]View)
	}
	a.views[state] = view
}

func (a *app) NavigateTo(state ViewState) {
	a.history = append(a.history, a.currentState)
	a.currentState = state
}

func (a *app) NavigateBack() {
	if len(a.history) == 0 {
		return
	}
	lastIndex := len(a.history) - 1
	a.currentState = a.history[lastIndex]
	a.history = a.history[:lastIndex]
}

func (a *app) Run() {
	var err error
	for a.currentState != ViewStateExit {
		deliveryutils.ClearScreen()
		view := a.views[a.currentState]

		nextState, renderErr := view.Render()
		if renderErr != nil {
			err = renderErr
			break
		}
		a.NavigateTo(nextState)
	}

	if err != nil {
		fmt.Printf("\nApplication exited with error: %v\n", err)
		fmt.Print("Press Enter to exit... ")
		bufio.NewReader(os.Stdin).ReadBytes('\n')
	}
}
