package main

import (
	"gophers.dev/cmds/consul-socket/internal/agent"
	"gophers.dev/cmds/consul-socket/internal/config"
)

func main() {
	args := config.ParseArguments()
	proxy := agent.New(args)
	proxy.Start()
}
