// Package moves implements a Moves API client.
package moves

import (
	"fmt"
	"net/http"
	"net/url"
	"time"
)

// Client is an HTTP client capable of making authenticated requests
// to the Moves API.
type Client struct {
	*http.Client        // Client is the http.Client used for requests.
	*Transport          // Transport is the http.RoundTripper used by the client.
	BaseURI      string // The base URI for requests.
}

// Weekday is like time.Weekday but starts at 1 for Sunday.
type Weekday int

// Time is a time.Time that can be JSON decoded by the RFC3339Short format.
type Time time.Time

// RFC3339Short defines the reference time for formatting.
const RFC3339Short = "20060102T150405Z0700"

const (
	Sunday Weekday = iota + 1
	Monday
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
)

// New creates a Moves client using the specified key and secret.
func New(key, secret string) *Client {
	transport := &Transport{
		Key:              key,
		Secret:           secret,
		AuthorizationURI: "https://api.moves-app.com/oauth/v1/authorize",
		ExchangeURI:      "https://api.moves-app.com/oauth/v1/access_token",
	}

	return &Client{
		Client:    &http.Client{Transport: transport},
		Transport: transport,
		BaseURI:   "https://api.moves-app.com/api/1.1",
	}
}

// get performs and HTTP get using the configuerd client.
func (c *Client) get(path string, query url.Values) (*http.Response, error) {
	rv, err := url.Parse(c.BaseURI + path)
	if err != nil {
		return nil, err
	}

	rv.RawQuery = query.Encode()

	return c.Get(rv.String())
}

// UnmarshalJSON implements json.Unmarshaler for Time.
func (t *Time) UnmarshalJSON(b []byte) error {
	if len(b) < 2 || b[0] != '"' || b[len(b)-1] != '"' {
		return fmt.Errorf("types: failed to unmarshal non-string value %q as an RFC 3339 time", b)
	}

	tm, err := time.Parse(RFC3339Short, string(b[1:len(b)-1]))
	if err != nil {
		return err
	}

	*t = Time(tm)

	return nil
}
