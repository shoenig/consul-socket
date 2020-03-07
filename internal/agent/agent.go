package agent

import (
	"log"
	"net/http"
	"os"

	"gophers.dev/cmds/consul-socket/internal/config"
	"gophers.dev/pkgs/loggy"
)

type Agent struct {
	socketPath string
	bindTo     string
	log        loggy.Logger
}

func New(c config.Arguments) *Agent {
	return &Agent{
		socketPath: c.SocketPath,
		bindTo:     c.BindAddress,
		log:        loggy.New("consul-socket"),
	}
}

func (a *Agent) Start() {
	a.log.Infof("starting up - bind to %s", a.bindTo)

	server := &http.Server{
		Addr:              a.bindTo,
		Handler:           newAPI(loggy.New("api")),
		ReadTimeout:       0,
		ReadHeaderTimeout: 0,
		WriteTimeout:      0,
		IdleTimeout:       0,
		MaxHeaderBytes:    0,
		TLSNextProto:      nil,
		ConnState:         nil,
		ErrorLog:          log.New(os.Stdout, "SERVER-ERROR", log.LstdFlags),
		BaseContext:       nil,
		ConnContext:       nil,
	}

	if err := server.ListenAndServe(); err != nil {
		panic(err)
	}
}
