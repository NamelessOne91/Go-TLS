package cmd

import (
	"github.com/spf13/cobra"
)

var createCmd = &cobra.Command{
	Use:   "create",
	Short: "create CA, certs or keys",
	Long:  "command to create resources (CA, certs, keys)",
}

func init() {
	rootCmd.AddCommand(createCmd)
}
