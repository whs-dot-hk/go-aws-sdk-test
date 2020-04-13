package cmd

import (
	"github.com/spf13/cobra"
)

var createCmd = &cobra.Command{
	Use:   "create [key or server]",
	Short: "Create key or server",
	Long:  `Create a key pair or a server`,
}

func init() {
	createCmd.AddCommand(createKeyCmd, createServerCmd)
}
