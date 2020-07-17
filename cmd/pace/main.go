package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
)

/*

	pace CardsService.CreateCard --Title="This is the title" --Body=""

*/

// ShortSHA is the git short hash at the time of this
// build.
var ShortSHA = `(missing)`

const (
	exitCodeErr       = 1
	exitCodeInterrupt = 2
)

func main() {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt)
	defer func() {
		signal.Stop(signalChan)
		cancel()
	}()
	go func() {
		select {
		case <-signalChan: // first signal, cancel context
			cancel()
		case <-ctx.Done():
		}
		<-signalChan // second signal, hard exit
		os.Exit(exitCodeInterrupt)
	}()
	if err := run(ctx, os.Args); err != nil {
		fmt.Fprintf(os.Stderr, "err: %s\n", err)
		os.Exit(exitCodeErr)
	}
}

func run(ctx context.Context, args []string) error {
	if err := runCommand(ctx, args); err != nil {
		return err
	}
	return nil
}
