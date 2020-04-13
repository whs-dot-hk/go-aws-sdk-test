package cmd

import (
	"github.com/whs-dot-hk/go-aws-sdk-test/server"

	"github.com/spf13/cobra"
)

var (
	createServerCmd = &cobra.Command{
		Use:   "server [name of the server]",
		Short: "Create server",
		Long:  `Create a server`,
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			server.CreateStack(args[0], keyName)
		},
	}
	keyName string
)

func init() {
	createServerCmd.Flags().StringVarP(&keyName, "key", "k", "", "key (required)")
	createServerCmd.MarkFlagRequired("key")
}
