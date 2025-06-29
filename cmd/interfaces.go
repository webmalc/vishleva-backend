package cmd

import "context"

// ErrorLogger logs errors.
type ErrorLogger interface {
	Error(args ...interface{})
}

// Runner runs the command.
type Runner interface {
	Run(names []string)
}

// ContextRunner runs the command with context.
type ContextRunner interface {
	Run(ctx context.Context, names []string)
}

// AdminCreator creates an admin user.
type AdminCreator interface {
	CreateAdmin(email, password string) error
}
