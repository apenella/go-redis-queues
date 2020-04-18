package publish

import (
	"context"

	"github.com/apenella/go-redis-queues/internal/application/transport"
	"github.com/apenella/go-redis-queues/internal/infrastructure/configuration"
	"github.com/apenella/go-redis-queues/pkg/command"

	"github.com/spf13/cobra"
)

type publishCmdFlags struct {
	Channel       string
	Message       string
	TransportType string
}

var publishCmdFlagsVar *publishCmdFlags

//  NewCommand return an stevedore command object for get
func NewCommand(ctx context.Context, config *configuration.Configuration) *command.AppCommand {

	publishCmdFlagsVar = &publishCmdFlags{}

	publishCmd := &cobra.Command{
		Use:   "publish",
		Short: "Publish message to Redis",
		Long:  "Publish message to Redis",
		RunE:  publishHandler(ctx, config),
	}

	command := &command.AppCommand{
		Command:       publishCmd,
		Configuration: config,
		Ctx:           ctx,
	}

	publishCmd.Flags().StringVarP(&publishCmdFlagsVar.Channel, "channel", "c", "", "Channel where the message will be published to.")
	publishCmd.MarkFlagRequired("channel")
	publishCmd.Flags().StringVarP(&publishCmdFlagsVar.Message, "message", "m", "", "Message to be published")
	publishCmd.MarkFlagRequired("message")
	publishCmd.Flags().StringVarP(&publishCmdFlagsVar.TransportType, "transport", "t", "", "Choose transport type [pubsub(default)|stream|fifo]")
	publishCmd.Flags().StringVarP(&publishCmdFlagsVar.Message, "group", "g", "default", "Consumer group used on stream transport")

	return command
}

func publishHandler(ctx context.Context, config *configuration.Configuration) command.CobraRunEFunc {
	return func(cmd *cobra.Command, args []string) error {

		if publishCmdFlagsVar.TransportType != "" {
			config.Transport = configuration.TransportToUint8(publishCmdFlagsVar.TransportType)
		}

		t, err := transport.NewTransport(config)
		if err != nil {
			return err
		}

		t.Publish(ctx, publishCmdFlagsVar.Channel, publishCmdFlagsVar.Message)
		return nil
	}
}
