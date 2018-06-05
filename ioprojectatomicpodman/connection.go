package ioprojectatomicpodman

import (
	"github.com/varlink/go/varlink"
)

var PodmanConnection *varlink.Connection

func NewPodmanConnection() (*varlink.Connection, error) {
	conn, err := varlink.NewConnection("tcp:10.70.48.27:8443")
	if err != nil {
		return nil, err
	}
	return conn, nil
}
