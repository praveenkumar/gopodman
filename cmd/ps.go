package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/docker/go-units"
	"github.com/praveenkumar/gopodman/ioprojectatomicpodman"
	"github.com/projectatomic/libpod/cmd/podman/formats"
	"github.com/urfave/cli"
	"reflect"
	"time"
)

var (
	psDescription = "List running containers"
	PsCommand     = cli.Command{
		Name:        "ps",
		Usage:       psDescription,
		Description: "List running containers",
		Action:      psCmd,
		ArgsUsage:   "",
	}
	psformat = "table {{.ID}}\t{{.Image}}\t{{.Command}}\t{{.Created}}\t{{.Status}}\t{{.Name}}\t"
)

type psTemplateParams struct {
	ID          string
	Image       string
	Command     string
	Created     string
	CreatedTime time.Time
	Status      string
	Name        string
}

type psSorted []psTemplateParams

// generatePsOutput generates the list of containers
func generatePsOutput(containers []ioprojectatomicpodman.ListContainerData) error {
	if len(containers) == 0 {
		return nil
	}
	var out formats.Writer
	psOutput := getPsTemplateOutput(containers)
	out = formats.StdoutTemplateArray{Output: psToGeneric(psOutput), Template: psformat, Fields: psOutput[0].headerMap()}

	return formats.Writer(out).Out()
}

// imagesToGeneric creates an empty array of interfaces for output
func psToGeneric(templParams []psTemplateParams) (genericParams []interface{}) {
	for _, v := range templParams {
		genericParams = append(genericParams, interface{}(v))
	}
	return
}

// getImagesTemplateOutput returns the images information to be printed in human readable format
func getPsTemplateOutput(containers []ioprojectatomicpodman.ListContainerData) (psOutput psSorted) {
	for _, container := range containers {
		// Todo: find a better solution for layout
		createdTime, _ := time.Parse("2008-01-02 15:04:05.000000000 +0530 IST", container.Createdat)
		params := psTemplateParams{
			ID:      container.Id,
			Image:   container.Image,
			Command: strings.Join(container.Command, " "),
			Created: units.HumanDuration(time.Since((createdTime))) + " ago",
			Status:  "Up " + container.Runningfor + " ago",
			Name:    container.Names,
		}
		if container.Containerrunning{
			psOutput = append(psOutput, params)
		}
	}

	return
}

// generate the accurate header based on template given
func (p *psTemplateParams) headerMap() map[string]string {
	v := reflect.Indirect(reflect.ValueOf(p))
	values := make(map[string]string)

	for i := 0; i < v.NumField(); i++ {
		key := v.Type().Field(i).Name
		value := key
		if value == "ID" {
			value = "Container" + value
		}
		values[key] = strings.ToUpper(value)
	}
	return values
}

func psCmd(c *cli.Context) {
	listContainerMethod := ioprojectatomicpodman.ListContainers()
	listContainerResponse, err := listContainerMethod.Call(ioprojectatomicpodman.PodmanConnection)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	generatePsOutput(listContainerResponse)
}
