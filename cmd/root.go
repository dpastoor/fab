package cmd

import (
	"fmt"
	"os"

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

var cfg settings

var rootCmd = &cobra.Command{
	Use:   "cmd",
	Short: "`fab`ricate new projects in a `fab`ulous way",
}

func Execute(version string, commit string, date string) {
	rootCmd.Version = version
	rootCmd.SetVersionTemplate("{{printf \"%s\\n\" .Version}}")
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.AddCommand(newDebugCmd())
	rootCmd.AddCommand(newGenerateCmd())
	rootCmd.AddCommand(newConfigCmd())

	// using viper so can take advantage of the casting and lookup capabilities of viper
	// even if don't need some of the more advanced functionality
	rootCmd.PersistentFlags().Bool("no-strict", false, "no strict mode")
	viper.BindPFlag("no-strict", rootCmd.PersistentFlags().Lookup("no-strict"))
	rootCmd.PersistentFlags().String("loglevel", "info", "log level")
	viper.BindPFlag("loglevel", rootCmd.PersistentFlags().Lookup("loglevel"))
	cobra.OnInitialize(initConfig)

}

func initConfig() {
	cfg.strict = !viper.GetBool("no-strict")
	cfg.loglevel = viper.GetString("loglevel")
	setLogLevel(cfg.loglevel)
}
