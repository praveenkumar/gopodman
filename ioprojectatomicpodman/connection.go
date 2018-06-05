package ioprojectatomicpodman

import (
	"github.com/varlink/go/varlink"
)

var PodmanConnection *varlink.Connection

func NewPodmanConnection(uri string) (*varlink.Connection, error) {
	conn, err := varlink.NewConnection(uri)
	if err != nil {
		return nil, err
	}
	return conn, nil
}
