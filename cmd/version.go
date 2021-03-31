package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

var (
	version string = "1.0"
	versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Check virtus version",
	Long: `Check what virtus version is you're machine runing
			Example: virtus version`,

	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("You're machine is runing the " + getVersion() + " virtus version")
	},
}
)

func getVersion() string{
	return version
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
