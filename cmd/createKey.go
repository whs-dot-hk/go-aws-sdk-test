package cmd

import (
	"github.com/whs-dot-hk/go-aws-sdk-test/key"

	"github.com/spf13/cobra"
)

var createKeyCmd = &cobra.Command{
	Use:   "key [name of the key]",
	Short: "Create key",
	Long:  "Create a key pair",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		key.CreateKey(args[0])
	},
}
