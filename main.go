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
		uri := c.GlobalString("podman-varlink-uri")
		if uri == "" {
			uri = os.Getenv("PODMAN_VARLINK_URI")
		}
		if uri == "" {
			return fmt.Errorf("Required podman varlink URL")
		}
		ioprojectatomicpodman.PodmanConnection, err = ioprojectatomicpodman.NewPodmanConnection(uri)
		if err != nil {
			return err
		}
		return nil
	}

	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "podman-varlink-uri",
			Usage: "Podman varlink URI (ex: tcp:127.0.0.1:12345)",
		},
	}
	if err := app.Run(os.Args); err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
	}
}
