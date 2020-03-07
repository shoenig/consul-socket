package agent

import (
	"net/http"
	"time"

	"gophers.dev/pkgs/loggy"
)

const (
	apiReadTimeout = 10 * time.Second
)

func newAPI(log loggy.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Tracef("%s request to %q", r.Method, r.URL.EscapedPath())
	}
}
