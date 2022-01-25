package cmd

import (
	"os"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

func configInit(_ *cobra.Command, args []string) {
	// this will return like:
	// fab config init
	// osx:
	// /Users/<user>/Library/Application Support/fab/config.yml
	// linux:
	// /home/<user>/.config/fab/config.yml
	cfgPath, err := getDefaultConfigPath()
	if err != nil {
		log.Fatal(err)
		return
	}
	_, err = os.Create(cfgPath)
	if err != nil {
		log.Fatal(err)
		return
	}
	log.Infof("configuration file created at: \"%s\"", cfgPath)
}

func newConfigInitCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "init",
		Short: "initialize a configuration file",
		Run:   configInit,
	}

	return cmd
}
