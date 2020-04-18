package get

import (
	"context"

	getconfiguration "github.com/apenella/go-redis-queues/internal/application/command/get/configuration"
	"github.com/apenella/go-redis-queues/internal/infrastructure/configuration"
	"github.com/apenella/go-redis-queues/pkg/command"

	"github.com/spf13/cobra"
)

//  NewCommand return an stevedore command object for get
func NewCommand(ctx context.Context, config *configuration.Configuration) *command.AppCommand {

	getCmd := &cobra.Command{
		Use:     "get",
		Aliases: []string{"list"},
		Short:   "Provide information related to go-redis-queues items",
		Long:    "",
		Run:     getHandler(),
	}

	command := &command.AppCommand{
		Command:       getCmd,
		Configuration: config,
		Ctx:           ctx,
	}

	command.AddCommand(getconfiguration.NewCommand(ctx, config))

	return command
}

func getHandler() command.CobraRunFunc {
	return func(cmd *cobra.Command, args []string) {
		cmd.HelpFunc()(cmd, args)
	}
}
