package cmd

import (
	"fmt"
	"encoding/json"
	"github.com/pkg/errors"
	"github.com/praveenkumar/gopodman/ioprojectatomicpodman"
	"github.com/urfave/cli"
)

var (
	inspectDescription = "Inspect an image or container"
	InspectCommand     = cli.Command{
		Name:        "inspect",
		Usage:       inspectDescription,
		Description: "List running containers",
		Action:      inspectCmd,
		ArgsUsage:   "",
	}
)

func inspectCmd(c *cli.Context) error {
	var inspectImageError, inspectContainerError error
	args := c.Args()
	if len(args) == 0 {
		return errors.Errorf("container or image name must be specified: gopodman inspect [options [...]] name")
	}
	name := args[0]
	imageInspectMethod := ioprojectatomicpodman.InspectImage()
	imageInspectResponse, err := imageInspectMethod.Call(ioprojectatomicpodman.PodmanConnection, name)
	if err != nil {
		inspectImageError = err
	}

	containerInspectMethod := ioprojectatomicpodman.InspectContainer()
	containerInspectResponse, err := containerInspectMethod.Call(ioprojectatomicpodman.PodmanConnection, name)
	if err != nil {
		inspectContainerError = err
	}

	if imageInspectResponse != "" {
		var raw map[string]interface{}
		err := json.Unmarshal([]byte(imageInspectResponse), &raw)
		if err != nil {
			return err
		}
		out, err := json.MarshalIndent(raw, "", " ")
		if err != nil {
			return err
		}
		fmt.Println(string(out))
	}

	if containerInspectResponse != "" {
		var raw map[string]interface{}
		err := json.Unmarshal([]byte(containerInspectResponse), &raw)
		if err != nil {
			return err
		}
		out, err := json.MarshalIndent(raw, "", "\t")
		if err != nil {
			return err
		}
		fmt.Println(out)
	}

	if inspectContainerError != nil && inspectImageError != nil {
		return errors.Errorf("Doesn't have any image or container as name: %s", name)
	}
	return nil
}
