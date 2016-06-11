package main

import (
	"fmt"
	"os"

	"github.com/umiyosh/anew/command"
	"github.com/urfave/cli"
)

var GlobalFlags = []cli.Flag{}

var Commands = []cli.Command{
	{
		Name:   "shell",
		Usage:  "anew shell hoge.sh",
		Action: command.CmdShell,
		Flags:  []cli.Flag{},
	},
	{
		Name:   "exec",
		Usage:  "anew exec 'make build && make test'",
		Action: command.CmdExec,
		Flags:  []cli.Flag{},
	},
}

func CommandNotFound(c *cli.Context, command string) {
	fmt.Fprintf(os.Stderr, "%s: '%s' is not a %s command. See '%s --help'.", c.App.Name, command, c.App.Name, c.App.Name)
	os.Exit(2)
}
