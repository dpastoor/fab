package cmd

import (
	"log"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

type settings struct {
	// strict mode will prevent the following:
	// - will check for path existence
	strict bool
	// logrus log level
	loglevel string
}

type rootCmd struct {
	cmd *cobra.Command
	cfg *settings
}

func Execute(version string, args []string) {
	newRootCmd(version).Execute(args)
}

func (cmd *rootCmd) Execute(args []string) {
	cmd.cmd.SetArgs(args)
	if err := cmd.cmd.Execute(); err != nil {
		log.Fatalf("failed with error: %s", err)
	}
}

func setGlobalSettings(cfg *settings) {
	cfg.strict = !viper.GetBool("no-strict")
	cfg.loglevel = viper.GetString("loglevel")
	setLogLevel(cfg.loglevel)
}
func newRootCmd(version string) *rootCmd {
	root := &rootCmd{cfg: &settings{}}
	cmd := &cobra.Command{
		Use:   "cmd",
		Short: "`fab`ricate new projects in a `fab`ulous way",
		PersistentPreRun: func(cmd *cobra.Command, args []string) {
			// need to set the config values here as the viper values
			// will not be processed until Execute, so can't
			// set them in the initializer.
			// If persistentPreRun is used elsewhere, should
			// remember to setGlobalSettings in the initializer
			setGlobalSettings(root.cfg)
		},
	}
	cmd.Version = version
	cmd.PersistentFlags().Bool("no-strict", false, "no strict mode")
	viper.BindPFlag("no-strict", cmd.PersistentFlags().Lookup("no-strict"))
	cmd.PersistentFlags().String("loglevel", "info", "log level")
	viper.BindPFlag("loglevel", cmd.PersistentFlags().Lookup("loglevel"))
	cmd.AddCommand(newDebugCmd(root.cfg))
	cmd.AddCommand(newGenerateCmd())
	cmd.AddCommand(newConfigCmd())
	cmd.AddCommand(newManCmd().cmd)

	root.cmd = cmd
	return root
}
