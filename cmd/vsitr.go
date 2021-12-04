package cmd

import (
	"github.com/spf13/cobra"
)

// vsitrCmd represents the vsitr command
var vsitrCmd = &cobra.Command{
	Use:   "vsitr",
	Short: "visitr a fast algorithm to delete a file",
	Long:  `change the byte contend of a file corresponding to the vsitr algorithm`,
	Run: func(cmd *cobra.Command, args []string) {
		runmethod(2, args)
	},
}

func init() {
	rootCmd.AddCommand(vsitrCmd)
}
