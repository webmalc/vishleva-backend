package cmd

import (
	"context"

	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

// CommandRouter is the main commands router.
type CommandRouter struct {
	logger          ErrorLogger
	rootCmd         *cobra.Command
	config          *Config
	serverRunner    ContextRunner
	bindatafsRunner Runner
}

// admin runs the server.
func (r *CommandRouter) server(_ *cobra.Command, args []string) {
	r.serverRunner.Run(context.Background(), args)
}

// bindatafs runs the bindatafs generator.
func (r *CommandRouter) bindatafs(_ *cobra.Command, args []string) {
	r.bindatafsRunner.Run(args)
}

// Run the router.
func (r *CommandRouter) Run() {
	r.rootCmd.AddCommand(
		&cobra.Command{
			Use:   "server",
			Short: "Run the server",
			Run:   r.server,
		},
		&cobra.Command{
			Use:   "bindatafs",
			Short: "Run the bindatafs generator",
			Run:   r.bindatafs,
		},
	)
	err := r.rootCmd.Execute()
	if err != nil {
		r.logger.Error(errors.Wrap(err, "root command"))
	}
}

// NewCommandRouter creates a new CommandRouter.
func NewCommandRouter(
	log ErrorLogger, server ContextRunner, bindata Runner,
) CommandRouter {
	config := NewConfig()

	return CommandRouter{
		config:          config,
		logger:          log,
		rootCmd:         &cobra.Command{Use: "vishleva_backend.app"},
		serverRunner:    server,
		bindatafsRunner: bindata,
	}
}
