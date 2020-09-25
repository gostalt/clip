package cmd

import (
	"fmt"
	"pwcli/password"

	"github.com/spf13/cobra"
)

func init() {
	generateCmd.Flags().IntVarP(&length, "length", "l", 12, "The length of the generated password")
	generateCmd.Flags().BoolVarP(&upper, "upper", "u", false, "Use uppercase letters in the generated password")
	generateCmd.Flags().BoolVarP(&digits, "digits", "d", false, "Use digits in the generated password")
	generateCmd.Flags().BoolVarP(&specials, "specials", "s", false, "Use special characters in the generated password")
}

var (
	length   int
	upper    bool
	digits   bool
	specials bool
)

var generateCmd = &cobra.Command{
	Use:     "generate [website]",
	Short:   "Generates a password",
	Example: "pwcli generate github.com --length 18 -dsu",
	Aliases: []string{"gen"},
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		var store password.PasswordStore
		store = password.NewPasswordFileStore()

		gp := password.Generate(password.GenerateOptions{
			Digits:   digits,
			Specials: specials,
			Upper:    upper,
			Length:   length,
		})

		store.Set(args[0], gp)

		fmt.Println(gp)
	},
}
