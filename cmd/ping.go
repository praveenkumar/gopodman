package cmd

import (
	"fmt"

	"github.com/praveenkumar/gopodman/ioprojectatomicpodman"
	"github.com/urfave/cli"
)

var (
	pingDescription = "display podman status"
	PingCommand     = cli.Command{
		Name:        "ping",
		Usage:       pingDescription,
		Description: `Information display here to make ensure podman is communicating`,
		Action:      pingCmd,
		ArgsUsage:   "",
	}
)

func pingCmd(c *cli.Context) error {
	pingMethod := ioprojectatomicpodman.Ping()
	pingResponse, err := pingMethod.Call(ioprojectatomicpodman.PodmanConnection)
	if err != nil {
		return err
	}
	fmt.Println(pingResponse.Message)
	return nil
}
