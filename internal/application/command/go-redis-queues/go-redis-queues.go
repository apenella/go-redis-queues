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

type redisPubSubCmdFlags struct {
	RedisHost     string
	RedisPort     uint16
	RedisDB       uint8
	RedisPassword string
}

var redisPubSubCmdFlagsVars *redisPubSubCmdFlags

func NewCommand(ctx context.Context, config *configuration.Configuration) *command.AppCommand {

	redisPubSubCmdFlagsVars = &redisPubSubCmdFlags{}

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
		PersistentPreRun: func(cmd *cobra.Command, args []string) {
			if redisPubSubCmdFlagsVars.RedisPort != configuration.RedisPortDefaultValue {
				config.RedisPort = redisPubSubCmdFlagsVars.RedisPort
			}

			if redisPubSubCmdFlagsVars.RedisHost != configuration.RedisHostDefaultValue {
				config.RedisHost = redisPubSubCmdFlagsVars.RedisHost
			}

			if redisPubSubCmdFlagsVars.RedisDB != configuration.RedisDBDefaultValue {
				config.RedisDB = redisPubSubCmdFlagsVars.RedisDB
			}

			if redisPubSubCmdFlagsVars.RedisPassword != configuration.RedisPasswordDefaultValue {
				config.RedisPassword = redisPubSubCmdFlagsVars.RedisPassword
			}
		},
	}

	appCommand := &command.AppCommand{
		Command:       redisPubSubCmd,
		Configuration: config,
		Ctx:           ctx,
	}

	redisPubSubCmd.PersistentFlags().StringVarP(&redisPubSubCmdFlagsVars.RedisHost, "host", "H", configuration.RedisHostDefaultValue, "Redis host")
	redisPubSubCmd.PersistentFlags().Uint16VarP(&redisPubSubCmdFlagsVars.RedisPort, "port", "P", configuration.RedisPortDefaultValue, "Redis port")
	redisPubSubCmd.PersistentFlags().Uint8VarP(&redisPubSubCmdFlagsVars.RedisDB, "database", "D", configuration.RedisDBDefaultValue, "Redis database")
	redisPubSubCmd.PersistentFlags().StringVarP(&redisPubSubCmdFlagsVars.RedisPassword, "password", "X", configuration.RedisPasswordDefaultValue, "Redis password")

	appCommand.AddCommand(middleware.New(middleware.PanicRecover).Then(middleware.ErrorManagement(consume.NewCommand(ctx, config))))
	appCommand.AddCommand(middleware.New(middleware.PanicRecover).Then(get.NewCommand(ctx, config)))
	appCommand.AddCommand(middleware.New(middleware.PanicRecover).Then(middleware.ErrorManagement(publish.NewCommand(ctx, config))))

	return appCommand
}

func redisPubSubHandler(cmd *cobra.Command, args []string) {
	cmd.HelpFunc()(cmd, args)
}
