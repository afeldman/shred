package cmd

import (
	"github.com/spf13/cobra"
)

// fastCmd represents the fast command
var fastCmd = &cobra.Command{
	Use:   "fast",
	Short: "delete a file fast",
	Long:  `Delete a file fast. only set 0 as element and delete the file`,
	Run: func(cmd *cobra.Command, args []string) {
		runmethod(1, args)
	},
}

func init() {

	rootCmd.AddCommand(fastCmd)
}
