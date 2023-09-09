package cmd

import (
	"fmt"
	"os"

	"github.com/NamelessOne91/go-tls/pkg/cert"
	"github.com/spf13/cobra"
)

var certKeyPath string
var certPath string
var certName string

func init() {
	createCmd.AddCommand(certCreateCmd)

	certCreateCmd.Flags().StringVarP(&certKeyPath, "key-out", "k", "server.key", "destination path for the server cert key")
	certCreateCmd.Flags().StringVarP(&certPath, "cert-out", "o", "server.crt", "destination path for the server cert")
	certCreateCmd.Flags().StringVarP(&certName, "name", "n", "", "name of the certificate in the config file")
	certCreateCmd.Flags().StringVar(&caKey, "ca-key", "ca.key", "CA key to sign the certificate")
	certCreateCmd.Flags().StringVar(&caCert, "ca-cert", "ca.crt", "CA cert for the certificate")

	certCreateCmd.MarkFlagRequired("ca-key")
	certCreateCmd.MarkFlagRequired("ca-cert")
	certCreateCmd.MarkFlagRequired("name")
}

var certCreateCmd = &cobra.Command{
	Use:   "cert",
	Short: "cert commands",
	Long:  "commands to create the certificates",
	Run: func(cmd *cobra.Command, args []string) {
		caKeyBytes, err := os.ReadFile(caKey)
		if err != nil {
			fmt.Printf("CA key read error :%v\n", err)
			return
		}

		caCertKeyBytes, err := os.ReadFile(caCert)
		if err != nil {
			fmt.Printf("CA cert read error :%v\n", err)
			return
		}

		certConfig, ok := config.Cert[certName]
		if !ok {
			fmt.Printf("Could not find certificate name (%s) in config\n", certName)
			return
		}

		err = cert.CreateCert(certConfig, caKeyBytes, caCertKeyBytes, certKeyPath, certPath)
		if err != nil {
			fmt.Printf("Create CA error: %v", err)
			return
		}
		fmt.Printf("Cert created. Key: %s, cert: %s\n", certKeyPath, certPath)
	},
}
