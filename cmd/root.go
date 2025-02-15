package cmd

import (
	"fmt"
	"os"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var (
	keep    bool
	report  bool
	verbose int

	// rootCmd represents the base command when called without any subcommands
	rootCmd = &cobra.Command{
		Use:   "shred",
		Short: "Delete files safely",
		Long: `Sometimes it is import to delete all information on the harddrive secure.
use this tool to delete the files as you wish`,
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) == 0 {
				log.Fatal("No files or directories provided. Use 'shred --help' for usage information.")
			}
			runmethod(3, args) // Default to usdod
		},
	}
)

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

}

func init() {
	rootCmd.PersistentFlags().BoolVarP(&keep, "keep", "k", false, "keep files")
	rootCmd.PersistentFlags().BoolVarP(&report, "report", "r", false, "Reporting about all operations?")
	rootCmd.PersistentFlags().IntVarP(&verbose, "verbose", "v", 4, "Verbose level (0-6)")
}
