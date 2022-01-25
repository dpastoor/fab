package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func debugCmd(_ *cobra.Command, args []string) {
	fmt.Printf("%#v\n", cfg)
}

func newDebugCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "debug",
		Short: "debug",
		Run:   debugCmd,
	}
	return cmd
}
