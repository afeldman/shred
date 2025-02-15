package cmd

import (
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// fastCmd represents the fast command
var fastCmd = &cobra.Command{
	Use:   "fast",
	Short: "Quickly overwrite files with zeroes (1 pass)",
	Long:  `The 'fast' method overwrites files with zeroes in a single pass, providing a quick but less secure deletion.`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			log.Fatal("No files or directories provided. Use 'shred --help' for usage information.")
		}
		err := runmethod(1, args)
		if err != nil {
			log.Fatalf("Error: %v", err)
		}
	},
}

func init() {
	rootCmd.AddCommand(fastCmd)
}
