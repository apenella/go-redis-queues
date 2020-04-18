package getconfiguration

import (
	"context"
	"fmt"

	"github.com/apenella/go-redis-queues/internal/infrastructure/configuration"
	"github.com/apenella/go-redis-queues/pkg/command"
	"github.com/spf13/cobra"
)

const (
	columnSeparator = " | "
)

//  NewCommand return an stevedore command object for get builders
func NewCommand(ctx context.Context, config *configuration.Configuration) *command.AppCommand {

	getConfigurationCmd := &cobra.Command{
		Use: "configuration",
		Aliases: []string{
			"config",
			"conf",
			"cfg",
		},
		Short: "Provide detail about go-redis-queues configuration",
		Long:  "Provide detail about go-redis-queues configuration",
		Run:   getConfigurationHandler(config),
	}

	command := &command.AppCommand{
		Command:       getConfigurationCmd,
		Configuration: config,
		Ctx:           ctx,
	}

	return command
}

func getConfigurationHandler(config *configuration.Configuration) command.CobraRunFunc {
	return func(cmd *cobra.Command, args []string) {
		fmt.Println(config.String())
	}
}
