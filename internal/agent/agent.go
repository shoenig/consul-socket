package agent

import (
	"log"
	"net/http"
	"os"
	"time"

	"gophers.dev/cmds/consul-socket/internal/config"
	"gophers.dev/pkgs/loggy"
)

const defaultTimeoutHTTP = 10 * time.Second

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
		ReadTimeout:       defaultTimeoutHTTP,
		ReadHeaderTimeout: defaultTimeoutHTTP,
		WriteTimeout:      defaultTimeoutHTTP,
		IdleTimeout:       defaultTimeoutHTTP,
		ErrorLog:          log.New(os.Stdout, "SERVER-ERROR", log.LstdFlags),
		BaseContext:       nil,
		ConnContext:       nil,
	}

	// todo: closer
	if err := server.ListenAndServe(); err != nil {
		panic(err)
	}
}
