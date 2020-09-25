package cmd

import "github.com/spf13/cobra"

var rootCmd = &cobra.Command{
	Use: "pwcli",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.UsageFunc()(cmd)
	},
}

// TODO: Add generate command, calls the add command with a random password, and returns it
func init() {
	rootCmd.AddCommand(listCmd, addCmd, getCmd, generateCmd)
}

func Execute() {
	rootCmd.Execute()
}
