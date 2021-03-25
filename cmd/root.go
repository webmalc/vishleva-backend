package cmd

import (
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

// CommandRouter is the main commands router.
type CommandRouter struct {
	logger          ErrorLogger
	rootCmd         *cobra.Command
	config          *Config
	adminRunner     Runner
	bindatafsRunner Runner
}

// admin runs admin server.
func (r *CommandRouter) admin(cmd *cobra.Command, args []string) {
	r.adminRunner.Run(args)
}

// bindatafs runs bindatafs generator.
func (r *CommandRouter) bindatafs(cmd *cobra.Command, args []string) {
	r.bindatafsRunner.Run(args)
}

// Run the router.
func (r *CommandRouter) Run() {
	r.rootCmd.AddCommand(
		&cobra.Command{
			Use:   "admin",
			Short: "Run the admin",
			Run:   r.admin,
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
	log ErrorLogger, admin, bindata Runner,
) CommandRouter {
	config := NewConfig()
	return CommandRouter{
		config:          config,
		logger:          log,
		rootCmd:         &cobra.Command{Use: "its-rankings.app"},
		adminRunner:     admin,
		bindatafsRunner: bindata,
	}
}
