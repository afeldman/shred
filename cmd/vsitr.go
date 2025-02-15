package cmd

import (
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// vsitrCmd represents the vsitr command
var vsitrCmd = &cobra.Command{
	Use:   "vsitr",
	Short: "visitr a fast algorithm to delete a file",
	Long:  `change the byte contend of a file corresponding to the vsitr algorithm`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			log.Fatal("No files or directories provided. Use 'shred --help' for usage information.")
		}
		err := runmethod(2, args)
		if err != nil {
			log.Fatalf("Error: %v", err)
		}
	},
}

func init() {
	rootCmd.AddCommand(vsitrCmd)
}
