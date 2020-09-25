package cmd

import (
	"fmt"
	"pwcli/password"

	"github.com/spf13/cobra"
)

var generateCmd = &cobra.Command{
	Use:   "generate",
	Short: "Generates a password",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		var store password.PasswordStore
		store = password.NewPasswordFileStore()

		gp := password.Generate(password.GenerateOptions{})

		store.Set(args[0], gp)

		fmt.Println(gp)
	},
}
