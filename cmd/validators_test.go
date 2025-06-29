package cmd

import (
	"testing"

	"github.com/spf13/cobra"
	"github.com/stretchr/testify/assert"
)

func Test_validateCreateAdminArgs(t *testing.T) {
	err := validateCreateAdminArgs(&cobra.Command{}, []string{
		"one", "two",
	})
	assert.NotNil(t, err)
	assert.Contains(t, err.Error(), "provide an email address")

	err = validateCreateAdminArgs(&cobra.Command{}, []string{
		"invalid email",
	})
	assert.NotNil(t, err)
	assert.Contains(t, err.Error(), "email is not valid")

	err = validateCreateAdminArgs(&cobra.Command{}, []string{
		"test@example.com",
	})
	assert.Nil(t, err)
}
