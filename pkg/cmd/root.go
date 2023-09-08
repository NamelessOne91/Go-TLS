package cmd

import (
	"fmt"
	"os"

	"github.com/NamelessOne91/go-tls/pkg/cert"
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v3"
)

type Config struct {
	CACert *cert.CACert          `yaml:"caCert"`
	Cert   map[string]*cert.Cert `yaml:"certs"`
}

var cfgFilePath string
var config Config

var rootCmd = &cobra.Command{
	Use:   "tls",
	Short: "tls is a CLI tool for TLS",
	Long:  "tls is a CLI tool for TLS. Mainly used for generation of X.509 certificates, can be extended.",
	Run: func(cmd *cobra.Command, args []string) {
		// Do Stuff Here
	},
}

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringVarP(&cfgFilePath, "config", "c", "", "config file (default is tls.yaml)")
}

func initConfig() {
	if cfgFilePath == "" {
		cfgFilePath = "tls.yml"
	}

	cfgFileBytes, err := os.ReadFile(cfgFilePath)
	if err != nil {
		fmt.Printf("Error reading config file: %v", err)
		return
	}

	err = yaml.Unmarshal(cfgFileBytes, &config)
	if err != nil {
		fmt.Printf("Error parsing config file: %v", err)
	}

}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
