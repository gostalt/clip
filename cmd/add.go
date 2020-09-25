package cmd

import (
	"fmt"
	"pwcli/password"

	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Adds a password",
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		var store password.PasswordStore
		store = password.NewPasswordFileStore()

		if err := store.Set(args[0], args[1]); err != nil {
			fmt.Println(err)
		}
	},
}
