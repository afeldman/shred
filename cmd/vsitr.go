package cmd

import (
	"github.com/spf13/cobra"
)

// vsitrCmd represents the vsitr command
var vsitrCmd = &cobra.Command{
	Use:   "vsitr",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		runmethod(2, args)
	},
}

func init() {
	rootCmd.AddCommand(vsitrCmd)
}
