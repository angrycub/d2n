package cmd

import (
	"fmt"
	"os"

	"github.com/angrycub/d2n/helper"
	"github.com/spf13/cobra"
)

// renderCmd represents the render command
var renderCmd = &cobra.Command{
	Use:   "render",
	Short: "Print job file to standard out",
	Long: `The render command allows you to incorporate the d2n application
in pipeline-driven workflow.

USAGE: d2n render [options] NAME IMAGE [command] [args]
	`,
	Args: cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		jobConfig.ParseArguments(args)

		out, err := jobConfig.RenderJob()
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}

		fmt.Println(out)
	},
}

func init() {
	rootCmd.AddCommand(renderCmd)
	helper.AddSharedDockerFlags(renderCmd, jobConfig)
}
