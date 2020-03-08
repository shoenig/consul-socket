package socket

import (
	"io"
	"net"
)

const (
	socketType = "unix"
)

type Forwarder struct {
	socketPath string
}

func New(path string) *Forwarder {
	return &Forwarder{
		socketPath: path,
	}
}

func (f *Forwarder) Start(source io.Reader) error {
	output, err := net.Dial(socketType, f.socketPath)
	if err != nil {
		return err
	}

	_ = output

	// io.Copy(output, source)
	return nil
}
