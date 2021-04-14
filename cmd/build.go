package cmd

import (
	"fmt"

	"github.com/angrycub/d2n/helper"
	"github.com/spf13/cobra"
)

// buildCmd represents the build command
var buildCmd = &cobra.Command{
	Use:   "build",
	Short: "Generate a Nomad job from a Docker command",
	Long:  `Writes a Nomad job file out as a file named «name».nomad.`,
	Args:  cobra.MinimumNArgs(3),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("build called")
	},
}

func init() {
	rootCmd.AddCommand(buildCmd)
	helper.AddSharedDockerFlags(buildCmd, jobConfig)
}
