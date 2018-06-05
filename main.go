package main

import (
	"fmt"
	"os"

	"github.com/praveenkumar/gopodman/cmd"
	"github.com/praveenkumar/gopodman/ioprojectatomicpodman"
	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "gopodman"
	app.Usage = "podman go client"
	app.Version = "0.0.1"

	app.Commands = []cli.Command{
		cmd.PingCommand,
		cmd.PodmanVersionCommand,
	}

	app.Before = func(c *cli.Context) error {
		var err error
		ioprojectatomicpodman.PodmanConnection, err = ioprojectatomicpodman.NewPodmanConnection()
		if err != nil {
			return err
		}
		return nil
	}
	if err := app.Run(os.Args); err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
	}
}
