package agent

import (
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"gophers.dev/cmds/consul-socket/internal/config"
	"gophers.dev/cmds/consul-socket/internal/socket"
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

	stopC := make(chan os.Signal, 1)
	signal.Notify(stopC, syscall.SIGTERM)

	server := &http.Server{
		Addr:              a.bindTo,
		Handler:           newAPI(socket.New(a.socketPath)),
		ReadTimeout:       defaultTimeoutHTTP,
		ReadHeaderTimeout: defaultTimeoutHTTP,
		WriteTimeout:      defaultTimeoutHTTP,
		IdleTimeout:       defaultTimeoutHTTP,
		ErrorLog:          log.New(os.Stdout, "SERVER-ERROR", log.LstdFlags),
		BaseContext:       nil,
		ConnContext:       nil,
	}

	go func(s *http.Server) {
		if err := s.ListenAndServe(); err != nil {
			a.log.Errorf("unable to listen and serve: %v", err)
			os.Exit(1)
		}
	}(server)

	select {
	case <-stopC:
		a.log.Infof("received SIGTERM, shutting down ...")
		if err := server.Close(); err != nil {
			a.log.Warnf("error shutting down server: %v", err)
		}
	}
}
