package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of TCurl",
	Long:  `All software has versions. This is TCurl's`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("TCurl Static Site Generator v0.1")
	},
}