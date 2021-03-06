package cmd

import (
	"fmt"
	"pwcli/password"

	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List stored passwords",
	Run: func(cmd *cobra.Command, args []string) {
		store := password.NewPasswordJsonStore()

		for i, service := range store.List() {
			fmt.Println(i, service)
		}
	},
}
