package cmd

import (
	"fmt"

	"github.com/asaskevich/govalidator"
	"github.com/spf13/cobra"
)

// validateCreateAdminArgs validate the create admin command args.
func validateCreateAdminArgs(_ *cobra.Command, args []string) error {
	if len(args) != 1 {
		return fmt.Errorf("provide an email address")
	}
	if !govalidator.IsEmail(args[0]) {
		return fmt.Errorf("email is not valid")
	}

	return nil
}
