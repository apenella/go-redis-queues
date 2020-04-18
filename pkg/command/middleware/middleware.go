package middleware

import (
	"context"
	"fmt"
	"os"

	"github.com/apenella/go-redis-queues/pkg/command"
	"github.com/spf13/cobra"
)

type MiddlewareFunc func(context.Context, command.CobraRunFunc) command.CobraRunFunc

type Chain []MiddlewareFunc

func (chain Chain) Then(c *command.AppCommand) *command.AppCommand {

	run := c.Command.Run
	for _, m := range chain {
		run = m(c.Ctx, run)
	}

	c.Command.Run = run
	return c
}

func New(middlewares ...MiddlewareFunc) Chain {
	c := Chain{}
	c = append(c, middlewares...)
	return c
}

func ErrorManagement(c *command.AppCommand) *command.AppCommand {

	if c.Command.RunE != nil {
		f := c.Command.RunE
		c.Command.Run = func(cmd *cobra.Command, args []string) {
			err := f(cmd, args)
			if err != nil {
				fmt.Println("Error: ", err.Error())
			}
		}
		c.Command.RunE = nil
	}

	return c
}

func PanicRecover(ctx context.Context, f command.CobraRunFunc) command.CobraRunFunc {
	return func(cmd *cobra.Command, args []string) {
		defer func() {
			if err := recover(); err != nil {
				msg := fmt.Sprintf("%s: %s", "Panic Recover", err)
				fmt.Fprintln(os.Stdout, msg)
			}
		}()
		f(cmd, args)
	}
}

// AuditCommand is a middleware which audits the command exection
func AuditCommand(ctx context.Context, f command.CobraRunFunc) command.CobraRunFunc {
	return func(cmd *cobra.Command, args []string) {

		fmt.Println("Executing command", cmd.Use, args, "'")
		f(cmd, args)
	}
}
