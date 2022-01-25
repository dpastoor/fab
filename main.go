package main

import "github.com/dpastoor/fab/cmd"

// https://goreleaser.com/cookbooks/using-main.version
var (
	version string = "dev"
	commit  string = "none"
	date    string = "unknown"
)

func main() {
	cmd.Execute(version, commit, date)
}
