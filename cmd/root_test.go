package cmd

import (
	"os"
	"testing"

	"github.com/spf13/cobra"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/webmalc/it-stats-rankings-scrapper/cmd/mocks"
	"github.com/webmalc/it-stats-rankings-scrapper/common/test"
)

// Should run the root command and log an error.
func TestCommandRouter_Run(t *testing.T) {
	m := &mocks.ErrorLogger{}
	r := &mocks.Runner{}
	cr := NewCommandRouter(m, r, r)
	os.Args = []string{"invalid", "invalid"}
	m.On("Error", mock.Anything).Return(nil).Once()
	cr.Run()
	m.AssertExpectations(t)
}

// Should create a command router object.
func TestNewCommandRouter(t *testing.T) {
	l := &mocks.ErrorLogger{}
	r := &mocks.Runner{}
	cr := NewCommandRouter(l, r, r)
	assert.Equal(t, l, cr.logger)
	assert.Equal(t, r, cr.adminRunner)
	assert.NotNil(t, cr.rootCmd)
}

func TestCommandRouter_admin(t *testing.T) {
	bindata := &mocks.Runner{}
	admin := &mocks.Runner{}
	cr := NewCommandRouter(&mocks.ErrorLogger{}, admin, bindata)
	admin.On("Run", mock.Anything).Return(nil).Once()
	cr.admin(&cobra.Command{}, []string{})
	admin.AssertExpectations(t)
}

func TestCommandRouter_bindatafs(t *testing.T) {
	bindata := &mocks.Runner{}
	admin := &mocks.Runner{}
	cr := NewCommandRouter(&mocks.ErrorLogger{}, admin, bindata)
	bindata.On("Run", mock.Anything).Return(nil).Once()
	cr.bindatafs(&cobra.Command{}, []string{})
	bindata.AssertExpectations(t)
}

// Setups the tests.
func TestMain(m *testing.M) {
	test.Run(m)
}
