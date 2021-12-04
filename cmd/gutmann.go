package cmd

import (
	"github.com/spf13/cobra"
)

// gutmannCmd represents the gutmann command
var gutmannCmd = &cobra.Command{
	Use:   "gutmann",
	Short: "really safe method. read gutmann paper",
	Long:  `delete data by changing the data in th file to a byte sequenze and delete the file `,
	Run: func(cmd *cobra.Command, args []string) {
		runmethod(4, args)
	},
}

func init() {
	rootCmd.AddCommand(gutmannCmd)
}
