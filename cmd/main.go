package main

import (
	"context"
	"main/di"
	"main/infra/program"
)

var (
	DistributionUse string = "false"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	cont := di.NewContainer(ctx, DistributionUse)
	prog := program.NewProgram(cont.RootVM)
	if err := prog.Start(); err != nil {
		panic(err)
	}
}
