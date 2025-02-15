package cmd

import (
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// usdodCmd represents the usdod command
var usdodCmd = &cobra.Command{
	Use:   "usdod",
	Short: "us department of devense method",
	Long:  `secure delete a file. this method was given out by the department of devense.`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			log.Fatal("No files or directories provided. Use 'shred --help' for usage information.")
		}
		err := runmethod(3, args)
		if err != nil {
			log.Fatalf("Error: %v", err)
		}
	},
}

func init() {
	rootCmd.AddCommand(usdodCmd)
}
