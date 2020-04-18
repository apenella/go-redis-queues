package pubsub

import (
	"context"

	"github.com/apenella/go-redis-queues/internal/application/command/consume"
	"github.com/apenella/go-redis-queues/internal/application/command/get"
	"github.com/apenella/go-redis-queues/internal/application/command/publish"
	"github.com/apenella/go-redis-queues/internal/infrastructure/configuration"
	"github.com/apenella/go-redis-queues/pkg/command"
	"github.com/apenella/go-redis-queues/pkg/command/middleware"
	"github.com/spf13/cobra"
)

func NewCommand(ctx context.Context, config *configuration.Configuration) *command.AppCommand {

	if ctx == nil {
		ctx = context.TODO()
	}

	redisPubSubCmd := &cobra.Command{
		Use:   "go-redis-queues",
		Short: "go-redis-queues interactuates with Redis and its types which could be used as queue",
		Long: `
go-redis-queues interactuates with Redis and its types which could be used as queue.
 These redis types are:
 - List
 - Pubsub
 - Stream
`,
		Run: redisPubSubHandler,
	}

	appCommand := &command.AppCommand{
		Command:       redisPubSubCmd,
		Configuration: config,
		Ctx:           ctx,
	}

	appCommand.AddCommand(middleware.New(middleware.PanicRecover).Then(middleware.ErrorManagement(consume.NewCommand(ctx, config))))
	appCommand.AddCommand(middleware.New(middleware.PanicRecover).Then(get.NewCommand(ctx, config)))
	appCommand.AddCommand(middleware.New(middleware.PanicRecover).Then(publish.NewCommand(ctx, config)))

	return appCommand
}

func redisPubSubHandler(cmd *cobra.Command, args []string) {
	cmd.HelpFunc()(cmd, args)
}
