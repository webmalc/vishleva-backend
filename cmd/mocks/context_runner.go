package mocks

import (
	"context"

	"github.com/stretchr/testify/mock"
)

// ContextRunner mocks the object
type ContextRunner struct {
	mock.Mock
}

// Run is method mock
func (r *ContextRunner) Run(ctx context.Context, names []string) {
	r.Called(names)
}
