consul-socket
=============

Forward `consul` HTTP requests over Unix socket

[![Go Report Card](https://goreportcard.com/badge/gophers.dev/cmds/consul-socket)](https://goreportcard.com/report/gophers.dev/cmds/consul-socket)
[![Build Status](https://travis-ci.com/shoenig/consul-socket.svg?branch=master)](https://travis-ci.com/shoenig/consul-socket)
[![GoDoc](https://godoc.org/gophers.dev/cmds/consul-socket?status.svg)](https://godoc.org/gophers.dev/cmds/consul-socket)
[![NetflixOSS Lifecycle](https://img.shields.io/osslifecycle/shoenig/consul-socket.svg)](OSSMETADATA)
[![GitHub](https://img.shields.io/github/license/shoenig/consul-socket.svg)](LICENSE)

# Project Overview

Module `gophers.dev/cmds/consul-socket` provides a simple, lightweight agent that
will proxy HTTP requests bound for Consul, and forward them over a Unix socket.

This is an **experimental** proof of concept for one possible way of enabling
Connect-native services work in `nomad`.

# Getting Started

The `consul-socket` command can be installed by running
```bash
$ go get gophers.dev/cmds/consul-socket
```

# Configuration

```bash
Usage of ./consul-socket:
  -bind string
    	bind address (i.e. inside network namespace) (default "localhost:8500")
  -socket string
    	filepath to consul unix socket (default "unix://var/run/consul.sock")
```

# Example Usage

Consul needs to be configured to listen to a unix socket for the `http` address.
See the `hack/consul.hcl` example file for a toy setup that enables 2 Connect-native
services to communicate with one another.

# Contributing

The `gophers.dev/cmds/consul-socket` module is always improving with new features
and error corrections. For contributing bug fixes and new features please file an issue.

# License

The `gophers.dev/cmds/consul-socket` module is open source under the [BSD-3-Clause](LICENSE) license.
