package socket

import (
	"context"
	"net"
	"net/http"
	"time"
)

const (
	socketType = "unix"

	socketTimeout = 1 * time.Hour
)

// New creates an HTTP client that makes requests to the specified socket,
// which is just a filepath.
func New(socket string) *http.Client {
	return &http.Client{
		Timeout: socketTimeout,
		Transport: &http.Transport{
			DialContext: func(_ context.Context, _, _ string) (net.Conn, error) {
				return net.Dial(socketType, socket)
			},
		},
	}
}
