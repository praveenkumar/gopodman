package cmd

import (
	"fmt"
	"os"

	"github.com/praveenkumar/gopodman/ioprojectatomicpodman"
	"github.com/urfave/cli"
)

var (
	podmanVersionDescription = "display podman version info"
	PodmanVersionCommand     = cli.Command{
		Name:        "podmanVersion",
		Usage:       podmanVersionDescription,
		Description: `Information display podman version`,
		Action:      podmanVersionCmd,
		ArgsUsage:   "",
	}
)

func podmanVersionCmd(c *cli.Context) {
	podmanVersionMethod := ioprojectatomicpodman.GetVersion()
	podmanVersionResponse, err := podmanVersionMethod.Call(ioprojectatomicpodman.PodmanConnection)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(podmanVersionResponse.Version)
}
