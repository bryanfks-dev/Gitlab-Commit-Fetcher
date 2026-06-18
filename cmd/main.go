package main

import (
	"bufio"
	"context"
	"fmt"
	"main/di"
	"os"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	cont := di.NewContainer(ctx)
	if err := cont.App.Run(); err != nil {
		fmt.Printf("\nApplication exited with error: %v\n", err)
		fmt.Print("Press Enter to exit... ")
		bufio.NewReader(os.Stdin).ReadBytes('\n')
	}
}
