package cmd

import (
	"fmt"
	"syscall"

	"golang.org/x/term"
)

// AdminCreatorRunner is the admin creator runner.
type AdminCreatorRunner struct {
	adminCreator   AdminCreator
	passwordReader func(int) ([]byte, error)
}

// Run runs the admin creator command.
func (s *AdminCreatorRunner) Run(args []string) {
	fmt.Println("Please enter your password:") // nolint
	result, err := s.passwordReader(syscall.Stdin)
	if err != nil {
		fmt.Println("error:", err) // nolint

		return
	}
	password := string(result)
	if password == "" {
		fmt.Println("password is empty") // nolint

		return
	}

	err = s.adminCreator.CreateAdmin(args[0], password)
	if err != nil {
		fmt.Println("error:", err) // nolint

		return
	}
	fmt.Println("admin has been created.") // nolint
}

func NewAdminCreatorRunner(adminCreator AdminCreator) *AdminCreatorRunner {
	return &AdminCreatorRunner{
		adminCreator:   adminCreator,
		passwordReader: term.ReadPassword,
	}
}
