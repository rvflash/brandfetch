package brandfetch

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net"
	"net/http"
	"net/url"
	"time"
)

// Configurator defines the interface used to define the settings.
type Configurator func(c *Client) error

// HTTPClient must be implemented by any HTTP client.
type HTTPClient interface {
	// Do sends an HTTP request and returns an HTTP response,
	// following policy (such as redirects, cookies, auth) as configured on the client.
	Do(req *http.Request) (*http.Response, error)
}

// SetHTTPClient overrides the default HTTP client used to perform customs API calls.
func SetHTTPClient(cli HTTPClient) Configurator {
	return func(c *Client) error {
		if cli == nil {
			return ErrHTTPClient
		}
		c.api = cli
		return nil
	}
}

// Connect creates the HTTP client to perform calls to the BrandfetchAPI.
func Connect(opts ...Configurator) (*Client, error) {
	var (
		c   = new(Client)
		err error
	)
	for _, opt := range append([]Configurator{SetHTTPClient(newHTTPClient())}, opts...) {
		err = opt(c)
		if err != nil {
			return nil, err
		}
	}
	return c, nil
}

const (
	apiURL         = "https://api.brandfetch.io/v2"
	searchEndpoint = "search"
)

// Client represents a Brandfetch Client API.
type Client struct {
	api HTTPClient
}

// BrandByName returns the "may be" brand behind this name using the following best effort mechanism: the first wins.
func (c *Client) BrandByName(ctx context.Context, name string) (*Brand, error) {
	res, err := c.BrandsByName(ctx, name)
	if err != nil {
		return nil, err
	}
	if len(res) > 0 {
		return &res[0], nil
	}
	return nil, ErrNoResults
}

// BrandsByName returns a list of maximum 3 brands matching the given name.
func (c *Client) BrandsByName(ctx context.Context, name string) ([]Brand, error) {
	err := c.ready(ctx)
	if err != nil {
		return nil, err
	}
	uri, err := url.JoinPath(apiURL, searchEndpoint, name)
	if err != nil {
		return nil, fmt.Errorf("%w: %s", ErrRequest, err)
	}
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return nil, fmt.Errorf("%w: %s", ErrRequest, err)
	}
	resp, err := c.api.Do(req)
	if err != nil {
		return nil, fmt.Errorf("%w: %s", ErrResponse, err)
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("%w: %s", ErrResponse, http.StatusText(resp.StatusCode))
	}
	buf, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("%w: %s", ErrResponse, err)
	}
	var res []Brand
	err = json.Unmarshal(buf, &res)
	if err != nil {
		return nil, fmt.Errorf("%w: %s", ErrResponse, err)
	}
	return res, nil
}

func (c *Client) ready(ctx context.Context) error {
	if c == nil || c.api == nil {
		return ErrHTTPClient
	}
	if ctx == nil {
		return context.Canceled
	}
	return nil
}

const (
	https = "https"
	// timeout is the default time limit for requests made by the HTTP Client.
	timeout = 10 * time.Second
	// keepAlive specifies the default interval between keep-alive probes for an active network connection.
	keepAlive = 600 * time.Second
	// MaxIdleConnections controls the maximum number of idle (keep-alive) connections across all hosts.
	maxIdleConnections = 10
	// maxIdleConnectionsPerHost controls the maximum idle (keep-alive) connections to keep per-host.
	maxIdleConnectionsPerHost = 10
)

func newHTTPClient() *http.Client {
	return &http.Client{
		Transport: &http.Transport{
			DialContext: (&net.Dialer{
				Timeout:   timeout,
				KeepAlive: keepAlive,
			}).DialContext,
			MaxIdleConns:        maxIdleConnections,
			MaxIdleConnsPerHost: maxIdleConnectionsPerHost,
		},
		Timeout: timeout,
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			if len(via) > 0 && via[0].URL.Scheme == https && req.URL.Scheme != https {
				lastURL := via[len(via)-1].URL
				return fmt.Errorf("redirected from secure rawURL %s to insecure rawURL %s", lastURL, req.URL)
			}
			return nil
		},
	}
}
