package cmd

import (
	"fmt"
	"pwcli/password"

	"github.com/spf13/cobra"
)

var getCmd = &cobra.Command{
	Use:   "get",
	Args:  cobra.ExactArgs(1),
	Short: "Gets a password",
	Run: func(cmd *cobra.Command, args []string) {
		var store password.PasswordStore
		store = password.NewPasswordFileStore()

		fmt.Println(store.Get(args[0]))
	},
}
