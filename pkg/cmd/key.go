package cmd

import (
	"fmt"

	"github.com/NamelessOne91/go-tls/pkg/key"
	"github.com/spf13/cobra"
)

var keyOut string
var keyLength int

func init() {
	createCmd.AddCommand(keyCreateCmd)

	keyCreateCmd.Flags().StringVarP(&keyOut, "key-out", "k", "key.pem", "destination path for key")
	keyCreateCmd.Flags().IntVarP(&keyLength, "key-length", "l", 4096, "destination path for key")
}

var keyCreateCmd = &cobra.Command{
	Use:   "key",
	Short: "key commands",
	Long:  "commands to create keys",
	Run: func(cmd *cobra.Command, args []string) {
		err := key.CreateRSAPrivateKeyAndSave(keyOut, keyLength)
		if err != nil {
			fmt.Printf("Create key error: %v", err)
			return
		}
		fmt.Printf("Key created %s with lenght %d\n", keyOut, keyLength)
	},
}
