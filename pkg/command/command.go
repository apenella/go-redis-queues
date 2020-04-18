package command

import (
	"context"
	"io"

	"github.com/apenella/go-redis-queues/internal/infrastructure/configuration"
	"github.com/spf13/cobra"
)

// CobraRunFunc is a cobra handler function
type CobraRunFunc func(cmd *cobra.Command, args []string)

// CobraRunEFunc is a cobra handler function which returns an error
type CobraRunEFunc func(cmd *cobra.Command, args []string) error

// AppCommand defines a stevedore command element
type AppCommand struct {
	Command       *cobra.Command
	Configuration *configuration.Configuration
	Ctx           context.Context
	Writer        io.Writer
}

// AddCommand method add a new subcommand to stevedore command
func (c *AppCommand) AddCommand(cmd *AppCommand) {
	c.Command.AddCommand(cmd.Command)
}

// Execute executes cobra command
func (c *AppCommand) Execute() error {
	if err := c.Command.Execute(); err != nil {
		return err
	}
	return nil
}
