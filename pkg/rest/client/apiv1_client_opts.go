package client

import (
	"net/http"
	"time"
)

// options is a struct that holds the options for the rest client
type options struct {
	transport http.RoundTripper
	timeout   time.Duration
}

type Option interface {
	apply(*options)
}

func getDefaultOptions() *options {
	return &options{
		timeout: 30 * time.Second,
	}
}

type transportOption struct {
	transport http.RoundTripper
}

func (t transportOption) apply(opts *options) {
	opts.transport = t.transport
}

// WithOptTransport sets the transport for the rest client.
// Transport specifies the mechanism by which individual
// HTTP requests are made.
// If nil, http.DefaultTransport is used.
func WithOptTransport(transport http.RoundTripper) Option {
	return transportOption{transport}
}
