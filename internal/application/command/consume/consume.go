package consume

import (
	"context"

	"github.com/apenella/go-redis-queues/internal/application/transport"
	"github.com/apenella/go-redis-queues/internal/infrastructure/configuration"
	"github.com/apenella/go-redis-queues/pkg/command"

	"github.com/spf13/cobra"
)

type consumeCmdFlags struct {
	Channel       string
	TransportType string
	Timeout       uint16
}

var consumeCmdFlagsVar *consumeCmdFlags

//  NewCommand return an stevedore command object for get
func NewCommand(ctx context.Context, config *configuration.Configuration) *command.AppCommand {

	consumeCmdFlagsVar = &consumeCmdFlags{}

	consumeCmd := &cobra.Command{
		Use:   "consume",
		Short: "Consume message from Redis",
		Long:  "Consume message from Redis",
		RunE:  consumeHandler(ctx, config),
	}

	command := &command.AppCommand{
		Command:       consumeCmd,
		Configuration: config,
		Ctx:           ctx,
	}

	consumeCmd.Flags().StringVarP(&consumeCmdFlagsVar.Channel, "channel", "c", "", "Channel where the message will be consumed from.")
	consumeCmd.MarkFlagRequired("channel")
	consumeCmd.Flags().StringVarP(&consumeCmdFlagsVar.TransportType, "transport", "t", "", "Choose transport type [pubsub(default)|stream|fifo]")
	consumeCmd.Flags().Uint16VarP(&consumeCmdFlagsVar.Timeout, "timeout", "T", uint16(5), "Timeout is the time (in seconds) to wait before close the consumer since last message received")

	return command
}

func consumeHandler(ctx context.Context, config *configuration.Configuration) command.CobraRunEFunc {
	return func(cmd *cobra.Command, args []string) error {

		if consumeCmdFlagsVar.TransportType != "" {
			config.Transport = configuration.TransportToUint8(consumeCmdFlagsVar.TransportType)
		}

		t, err := transport.NewTransport(config)
		if err != nil {
			return err
		}

		t.Consume(ctx, consumeCmdFlagsVar.Channel)
		return nil
	}
}
