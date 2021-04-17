package cmd

import (
	"os"
	"testing"

	"github.com/spf13/cobra"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/webmalc/vishleva-backend/cmd/mocks"
	"github.com/webmalc/vishleva-backend/common/test"
)

// Should run the root command and log an error.
func TestCommandRouter_Run(t *testing.T) {
	m := &mocks.ErrorLogger{}
	r := &mocks.Runner{}
	c := &mocks.ContextRunner{}
	cr := NewCommandRouter(m, c, r)
	os.Args = []string{"invalid", "invalid"}
	m.On("Error", mock.Anything).Return(nil).Once()
	cr.Run()
	m.AssertExpectations(t)
}

// Should create a command router object.
func TestNewCommandRouter(t *testing.T) {
	l := &mocks.ErrorLogger{}
	r := &mocks.Runner{}
	c := &mocks.ContextRunner{}
	cr := NewCommandRouter(l, c, r)
	assert.Equal(t, l, cr.logger)
	assert.Equal(t, c, cr.serverRunner)
	assert.Equal(t, r, cr.bindatafsRunner)
	assert.NotNil(t, cr.rootCmd)
}

func TestCommandRouter_server(t *testing.T) {
	bindata := &mocks.Runner{}
	server := &mocks.ContextRunner{}
	cr := NewCommandRouter(&mocks.ErrorLogger{}, server, bindata)
	server.On("Run", mock.Anything).Return(nil).Once()
	cr.server(&cobra.Command{}, []string{})
	server.AssertExpectations(t)
}

func TestCommandRouter_bindatafs(t *testing.T) {
	bindata := &mocks.Runner{}
	server := &mocks.ContextRunner{}
	cr := NewCommandRouter(&mocks.ErrorLogger{}, server, bindata)
	bindata.On("Run", mock.Anything).Return(nil).Once()
	cr.bindatafs(&cobra.Command{}, []string{})
	bindata.AssertExpectations(t)
}

// Setups the tests.
func TestMain(m *testing.M) {
	test.Run(m)
}
