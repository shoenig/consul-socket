package config

import (
	"flag"
)

type Arguments struct {
	SocketPath  string
	BindAddress string
}

func ParseArguments() Arguments {
	var socketPath string
	var bindAddr string
	flag.StringVar(
		&socketPath,
		"socket",
		"unix://var/run/consul.sock",
		"filepath to consul unix socket",
	)
	flag.StringVar(
		&bindAddr,
		"bind",
		"localhost:8500",
		"bind address (i.e. inside network namespace)",
	)
	flag.Parse()
	return Arguments{
		SocketPath:  socketPath,
		BindAddress: bindAddr,
	}
}
