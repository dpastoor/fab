package cmd

import (
	"bytes"
	"fmt"
	"io/ioutil"

	"github.com/dpastoor/fab/internal/config"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v3"
)

var collections []string

func configAdd(_ *cobra.Command, args []string) {
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

	cfgBytes, err := ioutil.ReadFile(cfgPath)
	if err != nil {
		log.Fatal(err)
		return
	}
	t := yaml.Node{}
	err = yaml.Unmarshal(cfgBytes, &t)
	newt, _ := config.AddPathsToCollections(t, collections, true, true)
	var b bytes.Buffer
	yamlEncoder := yaml.NewEncoder(&b)
	yamlEncoder.SetIndent(2)
	yamlEncoder.Encode(&newt)
	fmt.Println(string(b.Bytes()))
	err = ioutil.WriteFile(cfgPath, b.Bytes(), 0644)
}

func newConfigAddCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "add",
		Short: "add elements to a config",
		Run:   configAdd,
	}
	cmd.Flags().StringSliceVar(&collections, "collection", []string{}, "collection path to add")
	return cmd
}
