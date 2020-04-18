package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	pubsub "github.com/apenella/go-redis-queues/internal/application/command/go-redis-queues"
	"github.com/apenella/go-redis-queues/internal/infrastructure/configuration"
)

func main() {

	osSignalTerm := make(chan os.Signal)
	appDone := make(chan uint8)

	config := configuration.New()
	ctx, cancelFunc := context.WithCancel(context.Background())

	go func() {
		app := pubsub.NewCommand(ctx, config)
		app.Execute()
		appDone <- uint8(0)
	}()

	signal.Notify(osSignalTerm, syscall.SIGINT, syscall.SIGTERM)

	select {
	case <-osSignalTerm:
		fmt.Println("Process interrupted")
		cancelFunc()
		signal.Stop(osSignalTerm)
	case <-appDone:
	}

}
