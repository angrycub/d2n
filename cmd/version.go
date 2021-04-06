package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// versionCmd represents the version command
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of d2n",
	Long:  `Print the version number of d2n`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("d2n v0.9 -- HEAD")
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
