package cmd

import "context"

// ErrorLogger logs errors.
type ErrorLogger interface {
	Error(args ...interface{})
}

// Runner runs the command
type Runner interface {
	Run(names []string)
}
type ContextRunner interface {
	Run(ctx context.Context, names []string)
}
