package cmd

import (
	"github.com/spf13/cobra"
)

// gutmannCmd represents the gutmann command
var gutmannCmd = &cobra.Command{
	Use:   "gutmann",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		runmethod(4, args)
	},
}

func init() {
	rootCmd.AddCommand(gutmannCmd)
}
