package cmd

import (
	"github.com/spf13/cobra"
)

func configCmd(_ *cobra.Command, args []string) {
}

func newConfigCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "config",
		Short: "configuration",
		Run:   configCmd,
	}
	cmd.AddCommand(newConfigInitCmd())
	cmd.AddCommand(newConfigAddCmd())
	return cmd
}
