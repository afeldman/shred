package cmd

import (
	"fmt"
	"os"

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
		// Uncomment the following line if your bare application
		// has an action associated with it
	}
)

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	//cobra.CheckErr(rootCmd.Execute())
}

func init() {
	rootCmd.PersistentFlags().BoolVarP(&keep, "keep", "k", false, "keep files")
	rootCmd.PersistentFlags().BoolVarP(&report, "report", "r", false, "Reporting about all operations?")
	rootCmd.PersistentFlags().IntVarP(&verbose, "verbose", "v", 4, "verbose")
}
