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
	listImagesDescription = "List images in local storage"
	ListImagesCommand     = cli.Command{
		Name:        "images",
		Usage:       listImagesDescription,
		Description: "list images in local storage",
		Action:      listImagesCmd,
		ArgsUsage:   "",
	}
	imageformat = "table {{.Repository}}\t{{.Tag}}\t{{.ID}}\t{{.Created}}\t{{.Size}}\t"
)

type imagesTemplateParams struct {
	Repository  string
	Tag         string
	ID          string
	Created     string
	CreatedTime time.Time
	Size        string
}

type imagesSorted []imagesTemplateParams

// generateImagesOutput generates the images based on the format provided
func generateImagesOutput(images []ioprojectatomicpodman.ImageInList) error {
	if len(images) == 0 {
		return nil
	}
	var out formats.Writer
	imagesOutput := getImagesTemplateOutput(images)
	out = formats.StdoutTemplateArray{Output: imagesToGeneric(imagesOutput), Template: imageformat, Fields: imagesOutput[0].headerMap()}

	return formats.Writer(out).Out()
}

// imagesToGeneric creates an empty array of interfaces for output
func imagesToGeneric(templParams []imagesTemplateParams) (genericParams []interface{}) {
	for _, v := range templParams {
		genericParams = append(genericParams, interface{}(v))
	}
	return
}

// getImagesTemplateOutput returns the images information to be printed in human readable format
func getImagesTemplateOutput(images []ioprojectatomicpodman.ImageInList) (imagesOutput imagesSorted) {
	for _, img := range images {
		// Todo: find a better solution for layout
		createdTime, _ := time.Parse("2006-01-02 15:04:05.000000000 +0000 UTC", img.Created)
		imageID := img.Id
		// get all specified repo:tag pairs and print them separately
		for repo, tags := range reposToMap(img.RepoTags) {
			for _, tag := range tags {
				size := img.Size
				params := imagesTemplateParams{
					Repository:  repo,
					Tag:         tag,
					ID:          imageID,
					CreatedTime: createdTime,
					Created:     units.HumanDuration(time.Since((createdTime))) + " ago",
					Size:        units.HumanSizeWithPrecision(float64(size), 3),
				}
				imagesOutput = append(imagesOutput, params)
			}
		}
	}
	return
}

// HeaderMap produces a generic map of "headers" based on a line
// of output
func (i *imagesTemplateParams) headerMap() map[string]string {
	v := reflect.Indirect(reflect.ValueOf(i))
	values := make(map[string]string)

	for i := 0; i < v.NumField(); i++ {
		key := v.Type().Field(i).Name
		value := key
		if value == "ID" {
			value = "Image" + value
		}
		values[key] = strings.ToUpper(value)
	}
	return values
}

func listImagesCmd(c *cli.Context) {
	listImagesMethod := ioprojectatomicpodman.ListImages()
	listImagesResponse, err := listImagesMethod.Call(ioprojectatomicpodman.PodmanConnection)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	generateImagesOutput(listImagesResponse)
}
