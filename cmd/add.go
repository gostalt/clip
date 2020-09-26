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
		store := password.NewPasswordJsonStore()

		if err := store.Set(args[0], password.Account{"me@tomm.us", args[1]}); err != nil {
			fmt.Println(err)
		}
	},
}
