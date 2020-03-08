package agent

import (
	"bytes"
	"context"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"

	"github.com/pkg/errors"
	"gophers.dev/pkgs/loggy"
)

const (
	apiReadTimeout = 100 * time.Second
)

func newAPI(client *http.Client) http.HandlerFunc {
	log := loggy.New("api")

	return func(w http.ResponseWriter, r *http.Request) {
		log.Tracef("%s request to %q", r.Method, r.URL.EscapedPath())

		socketRequest, err := relayRequest(r)
		if err != nil {
			log.Errorf("failed to create forward request: %v", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		response, err := client.Do(socketRequest)
		if err != nil {
			log.Errorf("failed to execute request: %v", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		if response.StatusCode >= 400 {
			log.Errorf("failed request, code: %d, status: %s", response.StatusCode, response.Status)
			w.WriteHeader(response.StatusCode)
			return
		}

		bs, err := ioutil.ReadAll(response.Body)
		if err != nil {
			log.Errorf("failed to read request: %v", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		setHeaders(w, response.Header)
		if _, err = w.Write(bs); err != nil {
			log.Errorf("failed to write response: %v", err)
		}
	}
}

func setHeaders(w http.ResponseWriter, original http.Header) {
	for key, values := range original {
		duplicate := make([]string, len(values))
		copy(duplicate, values)
		w.Header()[key] = duplicate
	}
}

func relayRequest(original *http.Request) (*http.Request, error) {
	body, err := ioutil.ReadAll(original.Body)
	if err != nil {
		return nil, errors.Wrap(err, "unable to read original request body")
	}

	request := original.Clone(context.TODO())
	request.Body = ioutil.NopCloser(bytes.NewReader(body))
	request.Header.Set("X-Forwarded-For", original.RemoteAddr)
	request.Header.Set("User-Agent", "consul-socket (experimental)")
	request.RequestURI = ""
	request.URL = &url.URL{
		Scheme:     "http",
		Opaque:     "",
		User:       nil,
		Host:       "unix",
		Path:       original.URL.Path,
		RawPath:    "",
		ForceQuery: false,
		RawQuery:   original.URL.Query().Encode(),
		Fragment:   "",
	}

	return request, nil
}
