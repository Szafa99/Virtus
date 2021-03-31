
package cmd

import (
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "virtus",
	Short: "Starts a http server that servs a given html file",
	Long: `virtus is a command that makes it possible to serve a html file on a http server. 
	The server runs on port :8080`,
}

func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}

