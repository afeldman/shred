package cmd

import (
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// gutmannCmd represents the gutmann command
var gutmannCmd = &cobra.Command{
	Use:   "gutmann",
	Short: "really safe method. read gutmann paper",
	Long:  `delete data by changing the data in th file to a byte sequenze and delete the file `,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			log.Fatal("No files or directories provided. Use 'shred --help' for usage information.")
		}
		err := runmethod(4, args)
		if err != nil {
			log.Fatalf("Error: %v", err)
		}
	},
}

func init() {
	rootCmd.AddCommand(gutmannCmd)
}
