package cmd

import (
	"github.com/spf13/cobra"
)

// usdodCmd represents the usdod command
var usdodCmd = &cobra.Command{
	Use:   "usdod",
	Short: "us department of devense method",
	Long:  `secure delete a file. this method was given out by the department of devense.`,
	Run: func(cmd *cobra.Command, args []string) {
		runmethod(3, args)
	},
}

func init() {
	rootCmd.AddCommand(usdodCmd)
}
