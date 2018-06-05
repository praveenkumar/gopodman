package cmd

import (
	"fmt"
	"os"

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

func pingCmd(c *cli.Context) {
	pingMethod := ioprojectatomicpodman.Ping()
	pingResponse, err := pingMethod.Call(ioprojectatomicpodman.PodmanConnection)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(pingResponse.Message)
}
